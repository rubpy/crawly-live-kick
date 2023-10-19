package main

import (
	"context"
	"errors"
	"sync/atomic"
	"time"

	"github.com/eiannone/keyboard"

	"github.com/rubpy/crawly/csync"
)

//////////////////////////////////////////////////

type UIEventType uint

const (
	UIUnknownEventType UIEventType = iota

	UIKeyEventType
)

type UIEvent interface {
	Type() UIEventType
	StopPropagation()
}

type UIKeySubject struct {
	Key  keyboard.Key
	Rune rune
}

type UIKeyEvent struct {
	Subject  UIKeySubject
	KeyEvent keyboard.KeyEvent

	stopped atomic.Bool
}

func (e *UIKeyEvent) StopPropagation() {
	e.stopped.Store(true)
}

type UIBindHandler func(ui *UI, e *UIKeyEvent) error

type UI struct {
	keyBinds csync.Map[UIKeySubject, []UIBindHandler]

	listening atomic.Bool
	done      chan struct{}

	KeyDelay time.Duration
}

var (
	UIAlreadyListening = errors.New("already listening")
)

func NewUI() *UI {
	return &UI{
		done: make(chan struct{}),

		KeyDelay: 250 * time.Millisecond,
	}
}

func (ui *UI) Listening() bool {
	return ui.listening.Load()
}

func (ui *UI) Listen(ctx context.Context) (err error) {
	if ui.listening.Swap(true) {
		return UIAlreadyListening
	}
	defer ui.listening.Swap(false)

	done := ui.done

	keys, err := keyboard.GetKeys(1)
	if err != nil {
		return err
	}
	defer func() {
		_ = keyboard.Close()
	}()

	var lastKeyTimestamp time.Time
	keyDelay := ui.KeyDelay

listenLoop:
	for {
		select {
		case <-ctx.Done():
			break listenLoop
		case <-done:
			break listenLoop

		case ke := <-keys:
			{
				now := time.Now()
				if keyDelay > 0 && now.Sub(lastKeyTimestamp) < keyDelay {
					break
				}
				if ke.Err != nil {
					break
				}

				_, _ = ui.Trigger(ke)
				lastKeyTimestamp = now
			}
		}
	}

	return
}

func (ui *UI) Stop() {
	if !ui.listening.Load() {
		return
	}

	ui.done <- struct{}{}
}

func (ui *UI) Close() {
	ui.Stop()
}

func (ui *UI) Trigger(keyEvent keyboard.KeyEvent) (n int, err error) {
	subject := UIKeySubject{
		Rune: keyEvent.Rune,
	}
	if subject.Rune == 0 {
		subject.Key = keyEvent.Key
	}

	handlers, ok := ui.keyBinds.Load(subject)
	if !ok || len(handlers) == 0 {
		return
	}

	e := &UIKeyEvent{
		Subject:  subject,
		KeyEvent: keyEvent,
	}

	for _, h := range handlers {
		err = h(ui, e)
		n++

		if err != nil {
			break
		}

		if e.stopped.Load() {
			break
		}
	}

	return
}

func (ui *UI) BindKey(subject UIKeySubject, handler UIBindHandler) {
	if handler == nil {
		return
	}

	handlers, _ := ui.keyBinds.Load(subject)

	handlers = append(handlers, handler)
	ui.keyBinds.Store(subject, handlers)
}

func (ui *UI) Unbind(subject UIKeySubject) {
	ui.keyBinds.Delete(subject)
}

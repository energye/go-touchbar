package darwin

import (
	"encoding/json"
	"fmt"

	"github.com/LouisBrunner/go-touchbar/pkg/barbuilder"
)

type handlers struct {
	buttons      map[string]barbuilder.ButtonOnClick
	colorPickers map[string]barbuilder.ColorPickerOnSelected
	customs      map[string]barbuilder.CustomOnEvent
	pickers      map[string]barbuilder.PickerOnSelected
	scrubbers    map[string]barbuilder.ScrubberOnChange
	segments     map[string]barbuilder.SegmentedOnChange
	sliders      map[string]barbuilder.SliderOnChange
	steppers     map[string]barbuilder.StepperOnChange
}

const (
	eventButton      = "button"
	eventColorPicker = "color_picker"
	eventCustom      = "custom"
	eventPicker      = "picker"
	eventScrubber    = "scrubber"
	eventSegment     = "segment"
	eventSlider      = "slider"
	eventStepper     = "stepper"
)

type event struct {
	Kind   string
	Target string
	Data   json.RawMessage
}

func (me *touchBar) handleEventLogic(eventJSON string) error {
	event := event{}
	err := json.Unmarshal([]byte(eventJSON), &event)
	if err != nil {
		return err
	}

	switch event.Kind {
	case eventButton:
		handler, found := me.handlers.buttons[event.Target]
		if !found {
			return fmt.Errorf("unknown button %v", event.Target)
		}
		handler()

	case eventColorPicker:
		handler, found := me.handlers.colorPickers[event.Target]
		if !found {
			return fmt.Errorf("unknown color picker %v", event.Target)
		}
		data := barbuilder.ColorPickerColor{}
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return err
		}
		handler(data)

	case eventCustom:
		handler, found := me.handlers.customs[event.Target]
		if !found {
			return fmt.Errorf("unknown custom %v", event.Target)
		}
		data := barbuilder.CustomEvent{}
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return err
		}
		handler(data)

	case eventPicker:
		handler, found := me.handlers.pickers[event.Target]
		if !found {
			return fmt.Errorf("unknown picker %v", event.Target)
		}
		data := 0
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return err
		}
		handler(data)

	case eventScrubber:
		handler, found := me.handlers.scrubbers[event.Target]
		if !found {
			return fmt.Errorf("unknown scrubber %v", event.Target)
		}
		data := 0
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return err
		}
		handler(data)

	case eventSegment:
		handler, found := me.handlers.segments[event.Target]
		if !found {
			return fmt.Errorf("unknown segment %v", event.Target)
		}
		data := 0
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return err
		}
		handler(data)

	case eventSlider:
		handler, found := me.handlers.sliders[event.Target]
		if !found {
			return fmt.Errorf("unknown slider %v", event.Target)
		}
		data := float64(0)
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return err
		}
		handler(data)

	case eventStepper:
		handler, found := me.handlers.steppers[event.Target]
		if !found {
			return fmt.Errorf("unknown stepper %v", event.Target)
		}
		data := float64(0)
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return err
		}
		handler(data)

	default:
		return fmt.Errorf("unknown kind %v", event.Kind)
	}

	return nil
}

func (me *touchBar) handleEvent(eventJSON string) {
	err := me.handleEventLogic(eventJSON)
	if err != nil {
		// FIXME: no idea what to do, needs some kind of logger?
		fmt.Printf("event error: %v\n", err)
		return
	}
}

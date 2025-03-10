// Copyright 2013 The Ebitengine Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by genkeys.go using 'go generate'. DO NOT EDIT.

package ebiten

import (
	"fmt"
	"strings"

	"github.com/duplicants-ai/ebiten/internal/ui"
)

// A Key represents a keyboard key.
// These keys represent physical keys of US keyboard.
// For example, KeyQ represents Q key on US keyboards and ' (quote) key on Dvorak keyboards.
type Key int

// Keys.
const (
	KeyA              Key = Key(ui.KeyA)
	KeyB              Key = Key(ui.KeyB)
	KeyC              Key = Key(ui.KeyC)
	KeyD              Key = Key(ui.KeyD)
	KeyE              Key = Key(ui.KeyE)
	KeyF              Key = Key(ui.KeyF)
	KeyG              Key = Key(ui.KeyG)
	KeyH              Key = Key(ui.KeyH)
	KeyI              Key = Key(ui.KeyI)
	KeyJ              Key = Key(ui.KeyJ)
	KeyK              Key = Key(ui.KeyK)
	KeyL              Key = Key(ui.KeyL)
	KeyM              Key = Key(ui.KeyM)
	KeyN              Key = Key(ui.KeyN)
	KeyO              Key = Key(ui.KeyO)
	KeyP              Key = Key(ui.KeyP)
	KeyQ              Key = Key(ui.KeyQ)
	KeyR              Key = Key(ui.KeyR)
	KeyS              Key = Key(ui.KeyS)
	KeyT              Key = Key(ui.KeyT)
	KeyU              Key = Key(ui.KeyU)
	KeyV              Key = Key(ui.KeyV)
	KeyW              Key = Key(ui.KeyW)
	KeyX              Key = Key(ui.KeyX)
	KeyY              Key = Key(ui.KeyY)
	KeyZ              Key = Key(ui.KeyZ)
	KeyAltLeft        Key = Key(ui.KeyAltLeft)
	KeyAltRight       Key = Key(ui.KeyAltRight)
	KeyArrowDown      Key = Key(ui.KeyArrowDown)
	KeyArrowLeft      Key = Key(ui.KeyArrowLeft)
	KeyArrowRight     Key = Key(ui.KeyArrowRight)
	KeyArrowUp        Key = Key(ui.KeyArrowUp)
	KeyBackquote      Key = Key(ui.KeyBackquote)
	KeyBackslash      Key = Key(ui.KeyBackslash)
	KeyBackspace      Key = Key(ui.KeyBackspace)
	KeyBracketLeft    Key = Key(ui.KeyBracketLeft)
	KeyBracketRight   Key = Key(ui.KeyBracketRight)
	KeyCapsLock       Key = Key(ui.KeyCapsLock)
	KeyComma          Key = Key(ui.KeyComma)
	KeyContextMenu    Key = Key(ui.KeyContextMenu)
	KeyControlLeft    Key = Key(ui.KeyControlLeft)
	KeyControlRight   Key = Key(ui.KeyControlRight)
	KeyDelete         Key = Key(ui.KeyDelete)
	KeyDigit0         Key = Key(ui.KeyDigit0)
	KeyDigit1         Key = Key(ui.KeyDigit1)
	KeyDigit2         Key = Key(ui.KeyDigit2)
	KeyDigit3         Key = Key(ui.KeyDigit3)
	KeyDigit4         Key = Key(ui.KeyDigit4)
	KeyDigit5         Key = Key(ui.KeyDigit5)
	KeyDigit6         Key = Key(ui.KeyDigit6)
	KeyDigit7         Key = Key(ui.KeyDigit7)
	KeyDigit8         Key = Key(ui.KeyDigit8)
	KeyDigit9         Key = Key(ui.KeyDigit9)
	KeyEnd            Key = Key(ui.KeyEnd)
	KeyEnter          Key = Key(ui.KeyEnter)
	KeyEqual          Key = Key(ui.KeyEqual)
	KeyEscape         Key = Key(ui.KeyEscape)
	KeyF1             Key = Key(ui.KeyF1)
	KeyF2             Key = Key(ui.KeyF2)
	KeyF3             Key = Key(ui.KeyF3)
	KeyF4             Key = Key(ui.KeyF4)
	KeyF5             Key = Key(ui.KeyF5)
	KeyF6             Key = Key(ui.KeyF6)
	KeyF7             Key = Key(ui.KeyF7)
	KeyF8             Key = Key(ui.KeyF8)
	KeyF9             Key = Key(ui.KeyF9)
	KeyF10            Key = Key(ui.KeyF10)
	KeyF11            Key = Key(ui.KeyF11)
	KeyF12            Key = Key(ui.KeyF12)
	KeyF13            Key = Key(ui.KeyF13)
	KeyF14            Key = Key(ui.KeyF14)
	KeyF15            Key = Key(ui.KeyF15)
	KeyF16            Key = Key(ui.KeyF16)
	KeyF17            Key = Key(ui.KeyF17)
	KeyF18            Key = Key(ui.KeyF18)
	KeyF19            Key = Key(ui.KeyF19)
	KeyF20            Key = Key(ui.KeyF20)
	KeyF21            Key = Key(ui.KeyF21)
	KeyF22            Key = Key(ui.KeyF22)
	KeyF23            Key = Key(ui.KeyF23)
	KeyF24            Key = Key(ui.KeyF24)
	KeyHome           Key = Key(ui.KeyHome)
	KeyInsert         Key = Key(ui.KeyInsert)
	KeyIntlBackslash  Key = Key(ui.KeyIntlBackslash)
	KeyMetaLeft       Key = Key(ui.KeyMetaLeft)
	KeyMetaRight      Key = Key(ui.KeyMetaRight)
	KeyMinus          Key = Key(ui.KeyMinus)
	KeyNumLock        Key = Key(ui.KeyNumLock)
	KeyNumpad0        Key = Key(ui.KeyNumpad0)
	KeyNumpad1        Key = Key(ui.KeyNumpad1)
	KeyNumpad2        Key = Key(ui.KeyNumpad2)
	KeyNumpad3        Key = Key(ui.KeyNumpad3)
	KeyNumpad4        Key = Key(ui.KeyNumpad4)
	KeyNumpad5        Key = Key(ui.KeyNumpad5)
	KeyNumpad6        Key = Key(ui.KeyNumpad6)
	KeyNumpad7        Key = Key(ui.KeyNumpad7)
	KeyNumpad8        Key = Key(ui.KeyNumpad8)
	KeyNumpad9        Key = Key(ui.KeyNumpad9)
	KeyNumpadAdd      Key = Key(ui.KeyNumpadAdd)
	KeyNumpadDecimal  Key = Key(ui.KeyNumpadDecimal)
	KeyNumpadDivide   Key = Key(ui.KeyNumpadDivide)
	KeyNumpadEnter    Key = Key(ui.KeyNumpadEnter)
	KeyNumpadEqual    Key = Key(ui.KeyNumpadEqual)
	KeyNumpadMultiply Key = Key(ui.KeyNumpadMultiply)
	KeyNumpadSubtract Key = Key(ui.KeyNumpadSubtract)
	KeyPageDown       Key = Key(ui.KeyPageDown)
	KeyPageUp         Key = Key(ui.KeyPageUp)
	KeyPause          Key = Key(ui.KeyPause)
	KeyPeriod         Key = Key(ui.KeyPeriod)
	KeyPrintScreen    Key = Key(ui.KeyPrintScreen)
	KeyQuote          Key = Key(ui.KeyQuote)
	KeyScrollLock     Key = Key(ui.KeyScrollLock)
	KeySemicolon      Key = Key(ui.KeySemicolon)
	KeyShiftLeft      Key = Key(ui.KeyShiftLeft)
	KeyShiftRight     Key = Key(ui.KeyShiftRight)
	KeySlash          Key = Key(ui.KeySlash)
	KeySpace          Key = Key(ui.KeySpace)
	KeyTab            Key = Key(ui.KeyTab)
	KeyAlt            Key = Key(ui.KeyReserved0)
	KeyControl        Key = Key(ui.KeyReserved1)
	KeyShift          Key = Key(ui.KeyReserved2)
	KeyMeta           Key = Key(ui.KeyReserved3)
	KeyMax            Key = KeyMeta

	// Keys for backward compatibility.
	// Deprecated: as of v2.1.
	Key0            Key = Key(ui.KeyDigit0)
	Key1            Key = Key(ui.KeyDigit1)
	Key2            Key = Key(ui.KeyDigit2)
	Key3            Key = Key(ui.KeyDigit3)
	Key4            Key = Key(ui.KeyDigit4)
	Key5            Key = Key(ui.KeyDigit5)
	Key6            Key = Key(ui.KeyDigit6)
	Key7            Key = Key(ui.KeyDigit7)
	Key8            Key = Key(ui.KeyDigit8)
	Key9            Key = Key(ui.KeyDigit9)
	KeyApostrophe   Key = Key(ui.KeyQuote)
	KeyDown         Key = Key(ui.KeyArrowDown)
	KeyGraveAccent  Key = Key(ui.KeyBackquote)
	KeyKP0          Key = Key(ui.KeyNumpad0)
	KeyKP1          Key = Key(ui.KeyNumpad1)
	KeyKP2          Key = Key(ui.KeyNumpad2)
	KeyKP3          Key = Key(ui.KeyNumpad3)
	KeyKP4          Key = Key(ui.KeyNumpad4)
	KeyKP5          Key = Key(ui.KeyNumpad5)
	KeyKP6          Key = Key(ui.KeyNumpad6)
	KeyKP7          Key = Key(ui.KeyNumpad7)
	KeyKP8          Key = Key(ui.KeyNumpad8)
	KeyKP9          Key = Key(ui.KeyNumpad9)
	KeyKPAdd        Key = Key(ui.KeyNumpadAdd)
	KeyKPDecimal    Key = Key(ui.KeyNumpadDecimal)
	KeyKPDivide     Key = Key(ui.KeyNumpadDivide)
	KeyKPEnter      Key = Key(ui.KeyNumpadEnter)
	KeyKPEqual      Key = Key(ui.KeyNumpadEqual)
	KeyKPMultiply   Key = Key(ui.KeyNumpadMultiply)
	KeyKPSubtract   Key = Key(ui.KeyNumpadSubtract)
	KeyLeft         Key = Key(ui.KeyArrowLeft)
	KeyLeftBracket  Key = Key(ui.KeyBracketLeft)
	KeyMenu         Key = Key(ui.KeyContextMenu)
	KeyRight        Key = Key(ui.KeyArrowRight)
	KeyRightBracket Key = Key(ui.KeyBracketRight)
	KeyUp           Key = Key(ui.KeyArrowUp)
)

func (k Key) isValid() bool {
	switch k {
	case KeyA:
		return true
	case KeyB:
		return true
	case KeyC:
		return true
	case KeyD:
		return true
	case KeyE:
		return true
	case KeyF:
		return true
	case KeyG:
		return true
	case KeyH:
		return true
	case KeyI:
		return true
	case KeyJ:
		return true
	case KeyK:
		return true
	case KeyL:
		return true
	case KeyM:
		return true
	case KeyN:
		return true
	case KeyO:
		return true
	case KeyP:
		return true
	case KeyQ:
		return true
	case KeyR:
		return true
	case KeyS:
		return true
	case KeyT:
		return true
	case KeyU:
		return true
	case KeyV:
		return true
	case KeyW:
		return true
	case KeyX:
		return true
	case KeyY:
		return true
	case KeyZ:
		return true
	case KeyAlt:
		return true
	case KeyAltLeft:
		return true
	case KeyAltRight:
		return true
	case KeyArrowDown:
		return true
	case KeyArrowLeft:
		return true
	case KeyArrowRight:
		return true
	case KeyArrowUp:
		return true
	case KeyBackquote:
		return true
	case KeyBackslash:
		return true
	case KeyBackspace:
		return true
	case KeyBracketLeft:
		return true
	case KeyBracketRight:
		return true
	case KeyCapsLock:
		return true
	case KeyComma:
		return true
	case KeyContextMenu:
		return true
	case KeyControl:
		return true
	case KeyControlLeft:
		return true
	case KeyControlRight:
		return true
	case KeyDelete:
		return true
	case KeyDigit0:
		return true
	case KeyDigit1:
		return true
	case KeyDigit2:
		return true
	case KeyDigit3:
		return true
	case KeyDigit4:
		return true
	case KeyDigit5:
		return true
	case KeyDigit6:
		return true
	case KeyDigit7:
		return true
	case KeyDigit8:
		return true
	case KeyDigit9:
		return true
	case KeyEnd:
		return true
	case KeyEnter:
		return true
	case KeyEqual:
		return true
	case KeyEscape:
		return true
	case KeyF1:
		return true
	case KeyF2:
		return true
	case KeyF3:
		return true
	case KeyF4:
		return true
	case KeyF5:
		return true
	case KeyF6:
		return true
	case KeyF7:
		return true
	case KeyF8:
		return true
	case KeyF9:
		return true
	case KeyF10:
		return true
	case KeyF11:
		return true
	case KeyF12:
		return true
	case KeyF13:
		return true
	case KeyF14:
		return true
	case KeyF15:
		return true
	case KeyF16:
		return true
	case KeyF17:
		return true
	case KeyF18:
		return true
	case KeyF19:
		return true
	case KeyF20:
		return true
	case KeyF21:
		return true
	case KeyF22:
		return true
	case KeyF23:
		return true
	case KeyF24:
		return true
	case KeyHome:
		return true
	case KeyInsert:
		return true
	case KeyIntlBackslash:
		return true
	case KeyMeta:
		return true
	case KeyMetaLeft:
		return true
	case KeyMetaRight:
		return true
	case KeyMinus:
		return true
	case KeyNumLock:
		return true
	case KeyNumpad0:
		return true
	case KeyNumpad1:
		return true
	case KeyNumpad2:
		return true
	case KeyNumpad3:
		return true
	case KeyNumpad4:
		return true
	case KeyNumpad5:
		return true
	case KeyNumpad6:
		return true
	case KeyNumpad7:
		return true
	case KeyNumpad8:
		return true
	case KeyNumpad9:
		return true
	case KeyNumpadAdd:
		return true
	case KeyNumpadDecimal:
		return true
	case KeyNumpadDivide:
		return true
	case KeyNumpadEnter:
		return true
	case KeyNumpadEqual:
		return true
	case KeyNumpadMultiply:
		return true
	case KeyNumpadSubtract:
		return true
	case KeyPageDown:
		return true
	case KeyPageUp:
		return true
	case KeyPause:
		return true
	case KeyPeriod:
		return true
	case KeyPrintScreen:
		return true
	case KeyQuote:
		return true
	case KeyScrollLock:
		return true
	case KeySemicolon:
		return true
	case KeyShift:
		return true
	case KeyShiftLeft:
		return true
	case KeyShiftRight:
		return true
	case KeySlash:
		return true
	case KeySpace:
		return true
	case KeyTab:
		return true

	default:
		return false
	}
}

// String returns a string representing the key.
//
// If k is an undefined key, String returns an empty string.
func (k Key) String() string {
	switch k {
	case KeyA:
		return "A"
	case KeyB:
		return "B"
	case KeyC:
		return "C"
	case KeyD:
		return "D"
	case KeyE:
		return "E"
	case KeyF:
		return "F"
	case KeyG:
		return "G"
	case KeyH:
		return "H"
	case KeyI:
		return "I"
	case KeyJ:
		return "J"
	case KeyK:
		return "K"
	case KeyL:
		return "L"
	case KeyM:
		return "M"
	case KeyN:
		return "N"
	case KeyO:
		return "O"
	case KeyP:
		return "P"
	case KeyQ:
		return "Q"
	case KeyR:
		return "R"
	case KeyS:
		return "S"
	case KeyT:
		return "T"
	case KeyU:
		return "U"
	case KeyV:
		return "V"
	case KeyW:
		return "W"
	case KeyX:
		return "X"
	case KeyY:
		return "Y"
	case KeyZ:
		return "Z"
	case KeyAlt:
		return "Alt"
	case KeyAltLeft:
		return "AltLeft"
	case KeyAltRight:
		return "AltRight"
	case KeyArrowDown:
		return "ArrowDown"
	case KeyArrowLeft:
		return "ArrowLeft"
	case KeyArrowRight:
		return "ArrowRight"
	case KeyArrowUp:
		return "ArrowUp"
	case KeyBackquote:
		return "Backquote"
	case KeyBackslash:
		return "Backslash"
	case KeyBackspace:
		return "Backspace"
	case KeyBracketLeft:
		return "BracketLeft"
	case KeyBracketRight:
		return "BracketRight"
	case KeyCapsLock:
		return "CapsLock"
	case KeyComma:
		return "Comma"
	case KeyContextMenu:
		return "ContextMenu"
	case KeyControl:
		return "Control"
	case KeyControlLeft:
		return "ControlLeft"
	case KeyControlRight:
		return "ControlRight"
	case KeyDelete:
		return "Delete"
	case KeyDigit0:
		return "Digit0"
	case KeyDigit1:
		return "Digit1"
	case KeyDigit2:
		return "Digit2"
	case KeyDigit3:
		return "Digit3"
	case KeyDigit4:
		return "Digit4"
	case KeyDigit5:
		return "Digit5"
	case KeyDigit6:
		return "Digit6"
	case KeyDigit7:
		return "Digit7"
	case KeyDigit8:
		return "Digit8"
	case KeyDigit9:
		return "Digit9"
	case KeyEnd:
		return "End"
	case KeyEnter:
		return "Enter"
	case KeyEqual:
		return "Equal"
	case KeyEscape:
		return "Escape"
	case KeyF1:
		return "F1"
	case KeyF2:
		return "F2"
	case KeyF3:
		return "F3"
	case KeyF4:
		return "F4"
	case KeyF5:
		return "F5"
	case KeyF6:
		return "F6"
	case KeyF7:
		return "F7"
	case KeyF8:
		return "F8"
	case KeyF9:
		return "F9"
	case KeyF10:
		return "F10"
	case KeyF11:
		return "F11"
	case KeyF12:
		return "F12"
	case KeyF13:
		return "F13"
	case KeyF14:
		return "F14"
	case KeyF15:
		return "F15"
	case KeyF16:
		return "F16"
	case KeyF17:
		return "F17"
	case KeyF18:
		return "F18"
	case KeyF19:
		return "F19"
	case KeyF20:
		return "F20"
	case KeyF21:
		return "F21"
	case KeyF22:
		return "F22"
	case KeyF23:
		return "F23"
	case KeyF24:
		return "F24"
	case KeyHome:
		return "Home"
	case KeyInsert:
		return "Insert"
	case KeyIntlBackslash:
		return "IntlBackslash"
	case KeyMeta:
		return "Meta"
	case KeyMetaLeft:
		return "MetaLeft"
	case KeyMetaRight:
		return "MetaRight"
	case KeyMinus:
		return "Minus"
	case KeyNumLock:
		return "NumLock"
	case KeyNumpad0:
		return "Numpad0"
	case KeyNumpad1:
		return "Numpad1"
	case KeyNumpad2:
		return "Numpad2"
	case KeyNumpad3:
		return "Numpad3"
	case KeyNumpad4:
		return "Numpad4"
	case KeyNumpad5:
		return "Numpad5"
	case KeyNumpad6:
		return "Numpad6"
	case KeyNumpad7:
		return "Numpad7"
	case KeyNumpad8:
		return "Numpad8"
	case KeyNumpad9:
		return "Numpad9"
	case KeyNumpadAdd:
		return "NumpadAdd"
	case KeyNumpadDecimal:
		return "NumpadDecimal"
	case KeyNumpadDivide:
		return "NumpadDivide"
	case KeyNumpadEnter:
		return "NumpadEnter"
	case KeyNumpadEqual:
		return "NumpadEqual"
	case KeyNumpadMultiply:
		return "NumpadMultiply"
	case KeyNumpadSubtract:
		return "NumpadSubtract"
	case KeyPageDown:
		return "PageDown"
	case KeyPageUp:
		return "PageUp"
	case KeyPause:
		return "Pause"
	case KeyPeriod:
		return "Period"
	case KeyPrintScreen:
		return "PrintScreen"
	case KeyQuote:
		return "Quote"
	case KeyScrollLock:
		return "ScrollLock"
	case KeySemicolon:
		return "Semicolon"
	case KeyShift:
		return "Shift"
	case KeyShiftLeft:
		return "ShiftLeft"
	case KeyShiftRight:
		return "ShiftRight"
	case KeySlash:
		return "Slash"
	case KeySpace:
		return "Space"
	case KeyTab:
		return "Tab"
	}
	return ""
}

func keyNameToKeyCode(name string) (Key, bool) {
	switch strings.ToLower(name) {
	case "0":
		return Key0, true
	case "1":
		return Key1, true
	case "2":
		return Key2, true
	case "3":
		return Key3, true
	case "4":
		return Key4, true
	case "5":
		return Key5, true
	case "6":
		return Key6, true
	case "7":
		return Key7, true
	case "8":
		return Key8, true
	case "9":
		return Key9, true
	case "a":
		return KeyA, true
	case "b":
		return KeyB, true
	case "c":
		return KeyC, true
	case "d":
		return KeyD, true
	case "e":
		return KeyE, true
	case "f":
		return KeyF, true
	case "g":
		return KeyG, true
	case "h":
		return KeyH, true
	case "i":
		return KeyI, true
	case "j":
		return KeyJ, true
	case "k":
		return KeyK, true
	case "l":
		return KeyL, true
	case "m":
		return KeyM, true
	case "n":
		return KeyN, true
	case "o":
		return KeyO, true
	case "p":
		return KeyP, true
	case "q":
		return KeyQ, true
	case "r":
		return KeyR, true
	case "s":
		return KeyS, true
	case "t":
		return KeyT, true
	case "u":
		return KeyU, true
	case "v":
		return KeyV, true
	case "w":
		return KeyW, true
	case "x":
		return KeyX, true
	case "y":
		return KeyY, true
	case "z":
		return KeyZ, true
	case "alt":
		return KeyAlt, true
	case "altleft":
		return KeyAltLeft, true
	case "altright":
		return KeyAltRight, true
	case "apostrophe":
		return KeyApostrophe, true
	case "arrowdown":
		return KeyArrowDown, true
	case "arrowleft":
		return KeyArrowLeft, true
	case "arrowright":
		return KeyArrowRight, true
	case "arrowup":
		return KeyArrowUp, true
	case "backquote":
		return KeyBackquote, true
	case "backslash":
		return KeyBackslash, true
	case "backspace":
		return KeyBackspace, true
	case "bracketleft":
		return KeyBracketLeft, true
	case "bracketright":
		return KeyBracketRight, true
	case "capslock":
		return KeyCapsLock, true
	case "comma":
		return KeyComma, true
	case "contextmenu":
		return KeyContextMenu, true
	case "control":
		return KeyControl, true
	case "controlleft":
		return KeyControlLeft, true
	case "controlright":
		return KeyControlRight, true
	case "delete":
		return KeyDelete, true
	case "digit0":
		return KeyDigit0, true
	case "digit1":
		return KeyDigit1, true
	case "digit2":
		return KeyDigit2, true
	case "digit3":
		return KeyDigit3, true
	case "digit4":
		return KeyDigit4, true
	case "digit5":
		return KeyDigit5, true
	case "digit6":
		return KeyDigit6, true
	case "digit7":
		return KeyDigit7, true
	case "digit8":
		return KeyDigit8, true
	case "digit9":
		return KeyDigit9, true
	case "down":
		return KeyDown, true
	case "end":
		return KeyEnd, true
	case "enter":
		return KeyEnter, true
	case "equal":
		return KeyEqual, true
	case "escape":
		return KeyEscape, true
	case "f1":
		return KeyF1, true
	case "f2":
		return KeyF2, true
	case "f3":
		return KeyF3, true
	case "f4":
		return KeyF4, true
	case "f5":
		return KeyF5, true
	case "f6":
		return KeyF6, true
	case "f7":
		return KeyF7, true
	case "f8":
		return KeyF8, true
	case "f9":
		return KeyF9, true
	case "f10":
		return KeyF10, true
	case "f11":
		return KeyF11, true
	case "f12":
		return KeyF12, true
	case "f13":
		return KeyF13, true
	case "f14":
		return KeyF14, true
	case "f15":
		return KeyF15, true
	case "f16":
		return KeyF16, true
	case "f17":
		return KeyF17, true
	case "f18":
		return KeyF18, true
	case "f19":
		return KeyF19, true
	case "f20":
		return KeyF20, true
	case "f21":
		return KeyF21, true
	case "f22":
		return KeyF22, true
	case "f23":
		return KeyF23, true
	case "f24":
		return KeyF24, true
	case "graveaccent":
		return KeyGraveAccent, true
	case "home":
		return KeyHome, true
	case "insert":
		return KeyInsert, true
	case "intlbackslash":
		return KeyIntlBackslash, true
	case "kp0":
		return KeyKP0, true
	case "kp1":
		return KeyKP1, true
	case "kp2":
		return KeyKP2, true
	case "kp3":
		return KeyKP3, true
	case "kp4":
		return KeyKP4, true
	case "kp5":
		return KeyKP5, true
	case "kp6":
		return KeyKP6, true
	case "kp7":
		return KeyKP7, true
	case "kp8":
		return KeyKP8, true
	case "kp9":
		return KeyKP9, true
	case "kpadd":
		return KeyKPAdd, true
	case "kpdecimal":
		return KeyKPDecimal, true
	case "kpdivide":
		return KeyKPDivide, true
	case "kpenter":
		return KeyKPEnter, true
	case "kpequal":
		return KeyKPEqual, true
	case "kpmultiply":
		return KeyKPMultiply, true
	case "kpsubtract":
		return KeyKPSubtract, true
	case "left":
		return KeyLeft, true
	case "leftbracket":
		return KeyLeftBracket, true
	case "menu":
		return KeyMenu, true
	case "meta":
		return KeyMeta, true
	case "metaleft":
		return KeyMetaLeft, true
	case "metaright":
		return KeyMetaRight, true
	case "minus":
		return KeyMinus, true
	case "numlock":
		return KeyNumLock, true
	case "numpad0":
		return KeyNumpad0, true
	case "numpad1":
		return KeyNumpad1, true
	case "numpad2":
		return KeyNumpad2, true
	case "numpad3":
		return KeyNumpad3, true
	case "numpad4":
		return KeyNumpad4, true
	case "numpad5":
		return KeyNumpad5, true
	case "numpad6":
		return KeyNumpad6, true
	case "numpad7":
		return KeyNumpad7, true
	case "numpad8":
		return KeyNumpad8, true
	case "numpad9":
		return KeyNumpad9, true
	case "numpadadd":
		return KeyNumpadAdd, true
	case "numpaddecimal":
		return KeyNumpadDecimal, true
	case "numpaddivide":
		return KeyNumpadDivide, true
	case "numpadenter":
		return KeyNumpadEnter, true
	case "numpadequal":
		return KeyNumpadEqual, true
	case "numpadmultiply":
		return KeyNumpadMultiply, true
	case "numpadsubtract":
		return KeyNumpadSubtract, true
	case "pagedown":
		return KeyPageDown, true
	case "pageup":
		return KeyPageUp, true
	case "pause":
		return KeyPause, true
	case "period":
		return KeyPeriod, true
	case "printscreen":
		return KeyPrintScreen, true
	case "quote":
		return KeyQuote, true
	case "right":
		return KeyRight, true
	case "rightbracket":
		return KeyRightBracket, true
	case "scrolllock":
		return KeyScrollLock, true
	case "semicolon":
		return KeySemicolon, true
	case "shift":
		return KeyShift, true
	case "shiftleft":
		return KeyShiftLeft, true
	case "shiftright":
		return KeyShiftRight, true
	case "slash":
		return KeySlash, true
	case "space":
		return KeySpace, true
	case "tab":
		return KeyTab, true
	case "up":
		return KeyUp, true
	}
	return 0, false
}

// MarshalText implements encoding.TextMarshaler.
func (k Key) MarshalText() ([]byte, error) {
	return []byte(k.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler
func (k *Key) UnmarshalText(text []byte) error {
	key, ok := keyNameToKeyCode(string(text))
	if !ok {
		return fmt.Errorf("ebiten: unexpected key name: %s", string(text))
	}
	*k = key
	return nil
}

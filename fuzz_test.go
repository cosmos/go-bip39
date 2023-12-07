package bip39

import "testing"

func FuzzNewMnemonic(f *testing.F) {
	if testing.Short() {
		f.Skip("Running in -short mode")
	}

	// 1. Firstly add some seeds
	f.Add([]byte("00000000000000000000000000000000"))
	f.Add([]byte("7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f"))
	f.Add([]byte("80808080808080808080808080808080"))
	f.Add([]byte("18ab19a9f54a9274f03e5209a2ac8a91"))
	f.Add([]byte("b63a9c59a6e641f288ebc103017f1da9f8290b3da6bdef7b"))
	f.Add([]byte("18a2e1d81b8ecfb2a333adcb0c17a5b9eb76cc5d05db91a4"))
	f.Add([]byte("15da872c95a13dd738fbf50e427583ad61f18fd99f628c417a61cf8343c90419"))

	// 2. Now run the fuzzer.
	f.Fuzz(func(t *testing.T, entropy []byte) {
		_, _ = NewMnemonic(entropy)
	})
}

func FuzzMnemonicToByteArray(f *testing.F) {
	if testing.Short() {
		f.Skip("Running in -short mode")
	}

	// 1. Firstly add some seeds
	f.Add(" ")
	f.Add("")
	f.Add("abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon art art")
	f.Add("abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon art")
	f.Add("abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon agent")
	f.Add("abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon")
	f.Add("abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about")
	f.Add("abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon")
	f.Add("afford alter spike radar gate glance object seek swamp infant panel yellow")
	f.Add("beyond stage sleep clip because twist token leaf atom beauty genius food business side grid unable middle armed observe pair crouch tonight away coconut")
	f.Add("board blade invite damage undo sun mimic interest slam gaze truly inherit resist great inject rocket museum chief")
	f.Add("board flee heavy tunnel powder denial science ski answer betray cargo cat")
	f.Add("clutch control vehicle tonight unusual clog visa ice plunge glimpse recipe series open hour vintage deposit universe tip job dress radar refuse motion taste")
	f.Add("dignity pass list indicate nasty swamp pool script soccer toe leaf photo multiply desk host tomato cradle drill spread actor shine dismiss champion exotic")
	f.Add("dignity pass list indicate nasty")
	f.Add("exile ask congress lamp submit jacket era scheme attend cousin alcohol catch course end lucky hurt sentence oven short ball bird grab wing top")
	f.Add("indicate race push merry suffer human cruise dwarf pole review arch keep canvas theme poem divorce alter left")
	f.Add("jello better achieve collect unaware mountain thought cargo oxygen act hood bridge")
	f.Add("jelly better achieve collect unaware mountain thought cargo oxygen act hood bridge")
	f.Add("kiss carry display unusual confirm curtain upgrade antique rotate hello void custom frequent obey nut hole price segment")
	f.Add("legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth useful legal will will will")
	f.Add("legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth useful legal will")
	f.Add("legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth title")
	f.Add("legal winner thank year wave sausage worth useful legal winner thank yellow yellow")
	f.Add("legal winner thank year wave sausage worth useful legal winner thank yellow")
	f.Add("legal winner thank year wave sausage worth useful legal winner thanks year wave worth useful legal winner thank year wave sausage worth title")
	f.Add("letter advice cage absurd amount doctor acoustic avoid letter advice cage above")
	f.Add("letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic bless")
	f.Add("letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter always")
	f.Add("letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter always.")
	f.Add("letter advice cage absurd amount doctor acoustic avoid letter advice caged above")
	f.Add("letter advice cage absurd amount doctor acoustic avoid letters advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic bless")
	f.Add("renew stay biology evidence goat welcome casual join adapt armor shuffle fault little machine walk stumble urge swap")
	f.Add("turtle front uncle idea crush write shrug there lottery flower risk shell")
	f.Add("zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo wrong")
	f.Add("zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo when")
	f.Add("zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo why")
	f.Add("zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo vote")
	f.Add("zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo voted")

	// 2. Now run the fuzzer.
	f.Fuzz(func(t *testing.T, mnemonic string) {
		_, _ = MnemonicToByteArray(mnemonic)
	})
}

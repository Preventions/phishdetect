// PhishDetect
// Copyright (c) 2018-2020 Claudio Guarnieri.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package brand

// Apple brand properties.
func Apple() *Brand {
	name := "apple"
	original := []string{"apple", "appleid", "icloud", "itunes"}
	safelist := []string{
		"apple.com", "icloud.com", "mac.com", "airport.com", "applecomputer.com",
		"appleimac.com", "imac.com", "iphone.com", "iphone.org", "ipod.com",
		"itunes.com", "applemusic.com", "apple-dns.net", "cdn-apple.com",
	}
	suspicious := []string{
		"cpple", "epple", "ipple",
		"qpple", "aqple", "arple",
		"atple", "axple", "a0ple",
		"apqle", "aprle", "aptle",
		"apxle", "ap0le", "appme",
		"appne", "apphe", "appde",
		"appld", "applg", "appla",
		"applm", "applu", "xn--ppl-5qa00n",
		"xn--ppl-era09m", "xn--app1-epe", "xn--ppl-mra08m",
		"xn--ppl-dla7b", "xn--ppl-dla2c", "xn--ppl-dla7c",
		"xn--ple-tla0k", "xn--ple-tla014b", "xn--apl-2ra105a",
		"xn--apl-ura115a", "xn--al-3kaa0l", "xn--apl-mra125a",
		"xn--al-pbca05e", "xn--ppe-iwa676b", "xn--apl-5qa145a",
		"xn--ppl-8ka21a", "xn--ple-hla524b", "xn--al-gna67da",
		"xn--ale-poaa", "xn--ale-rbba", "xn--ppe-pla68b",
		"xn--apl-dma18h", "xn--apl-lma17h", "xn--apl-hma67h",
		"xn--ppe-gvb1227a", "xn--al-xma670aa", "xn--al-9ma060aa",
		"xn--al-3ma860aa", "xn--al-mna440aa", "xn--le-uia662aa",
		"xn--al-gna250aa", "xn--apl-6xc2d", "xn--ppl-tla7929a",
		"xn--ae-ehb10na", "xn--pp1e-foa", "xn--ppl-mnb724a",
		"xn--ae-3kaa45a", "xn--le-m1aa732a", "xn--ple-dla924b",
		"xn--le-m1aa24e", "xn--apl-dma683b", "xn--apl-hma183b",
		"xn--ape-iwa349a", "xn--ppl-hma59m", "xn--ppl-lma09m",
		"xn--ppl-dma00n", "xn--app-hvb46n", "xn--apl-lma773b",
		"xn--apl-tdd7c", "xn--ppe-iwa168a", "xn--apl-qoa0719a",
		"xn--ape-qoa93p", "xn--al-pbca59c", "xn--ple-zed8784a",
		"xn--pp1e-zsa", "xn--le-egca7325a", "xn--ppl-lma54i",
		"xn--pp1e-poa", "xn--app1-y4d", "xn--ppl-mra52g",
		"xn--ppl-era53g", "xn--ppl-5qa54g", "xn--ppl-rgzsf",
		"xn--apl-sbb5028a", "xn--ale-izc84d", "xn--apl-era135a",
		"xn--ppl-ura51g", "xn--ppl-gdb2318a", "xn--apl-6xc74a",
		"xn--le-w2a01ra", "xn--ppl-8ka2c", "xn--al-ija06fa",
		"xn--al-fja2fa", "xn--al-cja86fa", "xn--ppl-pla227a",
		"xn--al-olc6ba", "xn--ple-pla4k", "xn--appe-xpa",
		"xn--ppl-lma0419a", "xn--ap1e-wkb", "xn--apl-lma610b",
		"xn--le-lia872aa", "xn--ple-8oa55g", "xn--apie-bsa",
		"xn--le-kmca7784a", "xn--ape-gvb85u", "xn--ale-rbb913a",
		"xn--le-kmca74e", "xn--ppl-hma511c", "xn--ple-poax",
		"xn--ppl-dma021c", "xn--app1-e9d", "xn--ple-mnb56r",
		"xn--al-pbca57j", "xn--apl-poa1m", "xn--apl-poa1l",
		"xn--apl-poa1k", "xn--ple-pla977a", "xn--ple-fsb55p",
		"xn--apl-poa1i", "xn--ple-tla50i", "xn--apl-hma1h",
		"xn--apl-lma6g", "xn--apl-dma6h", "xn--ple-hla5l",
		"xn--ple-mnb93u", "xn--le-ila034aa", "xn--ppl-5qa019b",
		"xn--ppl-mra098b", "xn--ple-pla087a", "xn--ppl-2ra078b",
		"xn--ple-r5c80i", "xn--ppie-zna", "xn--ppie-9na",
		"xn--ple-dla42i", "xn--ple-poa202b", "xn--ale-poa429a",
		"xn--ple-hla91i", "xn--ple-r5c8325a", "xn--ale-s5ca",
		"xn--apl-sbb52w", "xn--ppl-tdd00f", "xn--le-m1aa23c",
		"xn--al-xma00ea", "xn--ppl-8ka2x", "xn--al-3kaa0i",
		"xn--ppl-8ka2y", "xn--al-9ma48da", "xn--ppl-8ka2z",
		"xn--al-3ma29da", "xn--ppl-8ka20a", "xn--ale-rbb48w",
		"xn--le-uia640aa", "xn--al-3kaa0m", "xn--ple-rbb250b",
		"xn--le-m1aa2808a", "xn--ple-fsb92s", "xn--ple-mnb04u",
		"xn--ppl-hma05i", "xn--ppl-dma55i", "xn--al-lmca04d",
		"xn--appi-ou5a", "xn--app-jwa942a", "xn--apl-s5c10h",
		"xn--ppl-2ra06m", "xn--ple-qbb33c", "xn--ple-0ed74e",
		"xn--ppl-tla753b", "xn--le-2kaa25h", "xn--le-ria072aa",
		"xn--ppl-era08k", "xn--apl-era661b", "xn--ple-pla01i",
		"xn--apl-s5c12a", "xn--apl-ura641b", "xn--apl-2ra631b",
		"xn--ppl-2ra05k", "xn--ppl-ura06k", "xn--ape-iwa249a",
		"xn--app-2ra39n", "xn--ple-lla024b", "xn--al-fja042aa",
		"xn--apl-s5c67b", "xn--al-cja442aa", "xn--apl-era16f",
		"xn--appi-opa", "xn--apl-mra15f", "xn--apl-5qa671b",
		"xn--apl-5qa17f", "xn--appi-jpa", "xn--ple-mnb46r",
		"xn--apl-ura14f", "xn--apl-2ra13f", "xn--ppl-5qa09k",
		"xn--al-gna043aa", "xn--apl-lma147a", "xn--apl-hma647a",
		"xn--apl-dma157a", "xn--ple-dla497a", "xn--le-2kaa710c",
		"xn--ppl-dla7w", "xn--ppl-dla7x", "xn--ppl-dla7y",
		"xn--ppl-dla7z", "xn--ppl-hma50p", "xn--ple-ooa80o",
		"xn--ale-rbb38w", "xn--ale-s5c37a", "xn--apl-1ed5h",
		"xn--ple-hzc38k", "xn--ppe-8oa61p", "xn--ae-3kaa93p",
		"xn--ppl-tdd0394a", "xn--ple-zed84e", "xn--ple-s5c70i",
		"xn--appe-i9t", "xn--apl-poa195a", "xn--ple-8oa426a",
		"xn--ppl-hla273b", "xn--ppl-8oa7f", "xn--apl-lma673b",
		"xn--ple-tla40i", "xn--ppl-8oa7h", "xn--appi-y4d",
		"xn--ppl-8oa7g", "xn--ppl-8oa7j", "xn--ppl-8oa7i",
		"xn--ppl-gdb291a", "xn--ple-rbb7e", "xn--le-oia450aa",
		"xn--appe-2h7o", "xn--al-lmca0005a", "xn--ple-gdb58y",
		"xn--apl-mra797a", "xn--apl-era708a", "xn--apl-5qa718a",
		"xn--ppl-mra570b", "xn--appi-e9d", "xn--ppl-ura560b",
		"xn--ppl-2ra550b", "xn--apl-2ra777a", "xn--apl-ura787a",
		"xn--ppl-pla2w", "xn--le-egca70i", "xn--ppl-pla2x",
		"xn--apl-lma27h", "xn--apl-hma77h", "xn--ppl-pla2v",
		"xn--ppl-pla2y", "xn--ppl-pla2z", "xn--ppe-mnb65a",
		"xn--ape-gvb89q", "xn--app-hma35b", "xn--ape-poa55a",
		"xn--app-dma85b", "xn--ppe-dla10c", "xn--ap1e-h6d",
		"xn--ap1e-6vd", "xn--ppl-gdb25v", "xn--apl-qoa019b",
		"xn--ppl-dhd5774a", "xn--le-ria2ia", "xn--a1e-poaa",
		"xn--ppl-hla725b", "xn--app-era32o", "xn--app-mra31o",
		"xn--al-ija425aa", "xn--ple-pla90i", "xn--app-5qa33o",
		"xn--ple-poa710c", "xn--ple-8ka570b", "xn--app-ura30o",
		"xn--le-iia282aa", "xn--ple-hzc86c", "xn--ppl-hla250c",
		"xn--al-ija632aa", "xn--apl-qoa583b", "xn--apie-vkb",
		"xn--ppl-dtd0524a", "xn--ple-fsb03s", "xn--le-obca28k",
		"xn--ple-lla060b", "xn--appi-epe", "xn--ple-8oa009a",
		"xn--pp1e-zna", "xn--ppie-4q5a", "xn--ple-r5c3z",
		"xn--appe-78b", "xn--al-n1aa562a", "xn--apl-ura215a",
		"xn--apl-2ra205a", "xn--apl-hma747a", "xn--al-p9b0ya",
		"xn--ppie-zmb", "xn--apl-5qa245a", "xn--ple-poa70o",
		"xn--apl-era235a", "xn--ple-ooa89l", "xn--apl-mra225a",
		"xn--ppl-fsb79o", "xn--ple-tla477a", "xn--appe-7pa",
		"xn--ap1e-cod", "xn--apl-1ed04d", "xn--ppe-lla67q",
		"xn--le-ria49fa", "xn--ple-lla587a", "xn--al-pbca5365a",
		"xn--apl-6xc1d", "xn--apl-qoa032b", "xn--le-iia075aa",
		"xn--ppl-lla225b", "xn--ppe-iwa11e", "xn--ale-poa946a",
		"xn--ppl-dla737a", "xn--ae-n1aa47f", "xn--apl-dma257a",
		"xn--apl-lma247a", "xn--ple-gdb01w", "xn--apl-s5c1635a",
		"xn--ale-s5c47a", "xn--apie-5vd", "xn--pp1e-p5b",
		"xn--apl-rbb62w", "xn--le-jbb26sa", "xn--appe-nza",
		"xn--le-ria855aa", "xn--al-3kaa0k", "xn--ppl-mnb70r",
		"xn--apl-6xc60e", "xn--appe-7md", "xn--apl-izc15e",
		"xn--ple-hla987a", "xn--ppl-mra0298a", "xn--ap1e-bsa",
		"xn--ple-lla0l", "xn--ple-dla52i", "xn--ple-s5c2z",
		"xn--ple-gdb48y", "xn--ple-mnb00y", "xn--apie-h6d",
		"xn--ppl-lla763b", "xn--ppl-hla2139a", "xn--ple-8oa45g",
		"xn--ple-s5c7325a", "xn--aie-izca", "xn--ple-0ed7784a",
		"xn--app-dma34q", "xn--ale-izc9y", "xn--al-mna86da",
		"xn--app-hma83q", "xn--ppl-lla7039a", "xn--ppl-tla730c",
		"xn--ple-izc76c", "xn--ple-hla097a", "xn--apie-wkb",
		"xn--le-6kc8da", "xn--pp1e-9na", "xn--app-ura8r",
		"xn--app-mra8s", "xn--apl-1ed0005a", "xn--app-5qa8u",
		"xn--le-m1aa250b", "xn--apl-2ra23f", "xn--le-lia0ja",
		"xn--apl-ura24f", "xn--apl-mra25f", "xn--ppl-2ra50g",
		"xn--apl-era26f", "xn--le-lia850aa", "xn--ppl-pla2a",
		"xn--ppl-lla7a", "xn--ppl-pla2b", "xn--ppl-pla7a",
		"xn--ale-izc8y", "xn--ple-dla070b", "xn--le-2kaa79l",
		"xn--ppl-8ka247a", "xn--al-3kaa095a", "xn--ple-8ka4m",
		"xn--al-n1aa52w", "xn--ppl-dma503b", "xn--ppl-lma592b",
		"xn--ppl-hma003b", "xn--apl-lma710b", "xn--ple-poa79l",
		"xn--apl-dma720b", "xn--ppl-lma00p", "xn--ple-qbb350b",
		"xn--ap1e-vkb", "xn--ale-rbb85z", "xn--ppl-8ka7c",
		"xn--le-obca76c", "xn--ple-lla950b", "xn--ppl-dma01p",
		"xn--ppl-hma5j", "xn--ppl-6xc52d", "xn--ppl-dma0k",
		"xn--ppe-iwa65i", "xn--ap1e-5vd", "xn--ppl-dla70a",
		"xn--ple-fsb98v", "xn--ape-gvb32o", "xn--ppl-lla727a",
		"xn--ple-izc2155a", "xn--ppl-lla740c", "xn--al-fja46fa",
		"xn--ple-poa25h", "xn--ple-8oa998a", "xn--ppl-6xc04l",
		"xn--ae-ehb98pa", "xn--ppl-8ka2d", "xn--appe-dnb",
		"xn--al-ija8ea", "xn--ple-ooa302b", "xn--ple-dla9l",
		"xn--apl-0ed14d", "xn--ap1e-g6d", "xn--ape-iwa285a",
		"xn--al-3kaa583b", "xn--ppl-8ka735b", "xn--al-n1aa024a",
		"xn--le-2kaax", "xn--ppl-lma011c", "xn--ple-hla424b",
		"xn--ppe-8ka60c", "xn--ale-izca", "xn--ple-dla597a",
		"xn--ple-lla51i", "xn--ppe-gvb191a", "xn--app-jwa9288a",
		"xn--le-uia8ha", "xn--ple-fsb09v", "xn--ple-izc28k",
		"xn--ple-lla41i", "xn--app1-opa", "xn--app1-jpa",
		"xn--app1-epa", "xn--ppe-gdb10f", "xn--le-2kaa70o",
		"xn--ppl-dhd54d", "xn--ppl-dla750c", "xn--pp1e-4na",
		"xn--ple-qbb3808a", "xn--ppe-8ka19q", "xn--ple-8ka03i",
		"xn--le-lia665aa", "xn--ppie-foa", "xn--al-cja235aa",
		"xn--al-fja825aa", "xn--le-obca2155a", "xn--a1e-izca",
		"xn--ple-tla9j", "xn--ppl-8ka283b", "xn--al-3kaa032b",
		"xn--pp1e-koa", "xn--apl-rbb649a", "xn--ple-ooay",
		"xn--ale-poa97g", "xn--le-xia055aa", "xn--appe-2ya",
		"xn--ppl-mnb74x", "xn--apl-ura687a", "xn--apl-tdd6c",
		"xn--apl-2ra677a", "xn--apl-era608a", "xn--apl-mra697a",
		"xn--apl-5qa618a", "xn--apl-rbb124a", "xn--aie-poaa",
		"xn--aie-rbba", "xn--ple-lla9k", "xn--ple-rbb2808a",
		"xn--ple-gdb90w", "xn--ple-8ka997a", "xn--ppe-dla68q",
		"xn--app-hvb4437a", "xn--al-fja020aa", "xn--al-ija610aa",
		"xn--al-cja420aa", "xn--ppl-8oa765a", "xn--ple-pla550b",
		"xn--ppl-dla235b", "xn--ale-izc94d", "xn--ppl-dtd0i",
		"xn--le-uia455aa", "xn--pp1e-pzb", "xn--apie-bod",
		"xn--ple-hla460b", "xn--apl-mra651b", "xn--ppl-dla773b",
		"xn--ppie-fse", "xn--ple-5cd4f", "xn--ple-tla904b",
		"xn--ape-iwa803a", "xn--ple-8ka92i", "xn--ppl-mnb20z",
		"xn--apl-0ed6h", "xn--ale-0eda", "xn--apie-6vd",
		"xn--ale-rbb813a", "xn--ppie-43d", "xn--apl-rbb662a",
		"xn--ple-hla02i", "xn--ple-rbb732a", "xn--app-lma33q",
		"xn--ale-poa87g", "xn--le-4eb60oa", "xn--al-lmca5h",
		"xn--ppl-8ka260c", "xn--al-3kaa019b", "xn--al-n1aa549a",
		"xn--ple-mnb99x", "xn--apie-g6d", "xn--al-fgca02a",
		"xn--le-4eb68la", "xn--ple-8ka008a", "xn--le-2kaa202b",
		"xn--ppl-fsb73v", "xn--ppie-poa", "xn--ple-dla0m",
		"xn--apl-dma28h", "xn--ple-8ka434b", "app1e",
		"xn--ple-8ka470b", "xn--ppl-gdb278a", "xn--apl-qoa095a",
		"xn--ppl-tla215b", "xn--ale-poa329a", "xn--ppl-fsb29w",
		"xn--ae-xqa96wa", "xn--ple-8oa062b", "xn--apl-5qa27f",
		"xn--ple-pla514b", "xn--apl-poa119b", "xn--al-mna42ya",
		"xn--al-gna23ya", "xn--ape-iwa385a", "xn--al-9ma04ya",
		"xn--le-uia09fa", "xn--al-3ma84ya", "xn--ppe-hla69b",
		"xn--al-xma65ya", "xn--le-iia4ja", "xn--apie-cod",
		"xn--app-jwa988a", "xn--apl-jzc5365a", "xn--ppl-tla717a",
		"xn--apl-dma620b", "xn--ppe-pla17q", "xn--al-fgca57b",
		"xn--apl-hma120b", "xn--ap1e-bod", "xn--ale-poa382b",
		"xn--app-2ra8q", "xn--ple-poa7409a", "xn--ple-tla577a",
		"xn--ppl-8oa7419a", "xn--ppl-5qa590b", "xn--al-xma463aa",
		"xn--ple-8ka534b", "xn--al-3ma653aa", "xn--al-9ma843aa",
		"xn--ppl-era580b", "xn--al-mna233aa", "xn--al-fgca00h",
		"xn--app1-jwa", "xn--app1-8va", "xn--app-era8t",
		"xn--le-m1aa7e", "xn--apl-6xc70e", "xn--ppie-p5b",
		"xn--ple-8oa952b", "xn--app1-eva", "xn--ple-hzc3155a",
		"xn--app1-yva", "xn--ppl-fsb712a", "xn--app1-ova",
		"xn--ppl-tla7x", "xn--apl-izc67j", "xn--ple-dla034b",
		"xn--aie-s5ca", "xn--apl-poa683b", "xn--le-ila241aa",
		"xn--ape-qoa45a", "xn--ple-lla914b", "xn--ple-ooa810c",
		"xn--ple-pla450b", "xn--app1-ou5a", "xn--ppl-gdb743a",
		"xn--ppl-6xc0755a", "xn--apl-poa132b", "xn--ppl-hla237a",
		"xn--le-w2a81wa", "xn--apl-poa1j", "xn--appi-omd",
		"xn--ppl-5cd2805a", "xn--apl-t5c0635a", "xn--ppe-tla18b",
		"xn--apie-csa", "xn--ppie-4na", "xn--ae-ehb18ka",
		"xn--ppl-pla2039a", "xn--ape-iwa73d", "xn--le-iia260aa",
		"xn--apl-sbb549a", "xn--ple-8oa526a", "xn--apl-0ed1005a",
		"xn--le-oia6ia", "xn--ae-xqa771aa", "xn--ppl-mnb7867a",
		"xn--le-xia262aa", "xn--ple-tla050b", "xn--ple-5cd5f",
		"xn--le-jbb47pa", "xn--apl-hma2h", "xn--apl-lma7g",
		"xn--apl-dma7h", "xn--ae-xqa31ca", "xn--ple-gdb542a",
		"xn--le-oia89fa", "xn--app-jwa965b", "xn--ple-lla487a",
		"xn--ppl-lla7v", "xn--ape-sbb47f", "xn--ppl-lla7x",
		"xn--ppl-lla7w", "xn--ppl-lla7z", "xn--ppl-lla7y",
		"xn--le-jbb45na", "xn--ppl-dla7139a", "xn--le-2kaa7409a",
		"xn--ape-iwa703a", "xn--le-xia68fa", "xn--apl-6xc64a",
		"xn--ppl-tlax", "xn--ppl-tla2a", "xn--apl-hma220b",
		"xn--pp1e-fse", "xn--ple-pla5k", "xn--ple-pla414b",
		"xn--ple-dla960b", "xn--ppe-lla19b", "xn--apl-rbb6028a",
		"xn--le-oia472aa", "xn--le-w2a03ta", "xn--apl-sbb562a",
		"xn--ppl-era009b", "xn--le-ila66ea", "xn--ple-rbb23c",
		"xn--ple-ooa35h", "xn--ppl-tla7a", "xn--ppl-ura088b",
		"xn--ppe-tla66q", "xn--ple-tla940b", "xn--le-ria050aa",
		"xn--apl-dma783b", "xn--apl-hma283b", "xn--ale-poa482b",
		"xn--ppl-hla2c", "xn--ppl-hla7b", "xn--ppl-hla2b",
		"xn--le-xia240aa", "xn--ppl-dma0519a", "xn--ppl-hma5419a",
		"xn--ple-ooa8409a", "xn--al-cja6fa", "xn--app-jwa440b",
		"xn--ape-iwa83d", "xn--ple-qbb8e", "xn--apl-sbb024a",
		"xn--ap1e-csa", "xn--ppie-koa", "xn--ple-qbb832a",
		"xn--ape-rbb57f", "xn--ppl-8ka2239a", "xn--ple-hla560b",
		"xn--appi-epa", "xn--al-n1aa5028a", "xn--le-iia60ga",
		"xn--le-4eb49qa", "xn--appi-eva", "xn--le-xia4ha",
		"xn--apl-t5c57b", "xn--appi-yva", "xn--ppl-5cd7p",
		"xn--appi-ova", "xn--ppl-8oa702b", "xn--appi-jwa",
		"xn--ale-poa846a", "xn--appi-8va", "xn--ape-gvb79q",
		"xn--ppl-pla715b", "xn--ppl-lma0j", "xn--ppie-pzb",
		"xn--ppe-fsb6k", "xn--ape-gvb22o", "xn--app-hvb480a",
		"xn--ppe-iwa6078a", "xn--ale-rbb95z", "xn--ppl-hla20a",
		"xn--ppl-hla2z", "xn--ppl-5qa0498a", "xn--ppl-era0398a",
		"xn--al-3kaa0719a", "xn--ppl-hla2w", "xn--ppl-ura0198a",
		"xn--ape-gvb75u", "xn--ppl-2ra0098a", "xn--ppl-hla2y",
		"xn--ppl-hla2x", "xn--apl-5qa771b", "xn--apl-era761b",
		"xn--apl-mra751b", "xn--apl-ura741b", "xn--apl-2ra731b",
		"xn--app1-omd", "xn--a1e-s5ca", "xn--ple-8ka5m",
		"xn--le-ila22za", "xn--apl-poa1719a", "xn--ape-poa04p",
		"xn--apl-izc69c", "xn--le-egca2z", "xn--apl-t5c02a",
		"xn--pp1e-43d", "xn--ppl-5cd2a", "xn--ppl-8oa263b",
		"xn--app-lma84b", "xn--ppl-pla263b", "xn--ppe-gvb67t",
		"xn--apl-jzc57j", "xn--al-p9b86ca", "xn--ppe-8oa13a",
		"xn--pp1e-zmb", "xn--ppie-zsa", "xn--ae-xqa98ya",
		"xn--ple-rbb24e", "xn--ppe-hla18q", "xn--ple-hla4l",
		"xn--ppl-8oa788b", "xn--ppl-tla7v", "xn--ppl-tla7w",
		"xn--ppl-tla7u", "xn--al-fgca0635a", "xn--apl-jzc05e",
		"xn--pp1e-4q5a", "xn--ppl-tla7y", "xn--app-hvb40u",
		"xn--ppe-iwa66k", "xn--aie-0eda", "xn--le-lia20ga",
		"xn--ppl-fsb7747a", "xn--le-oia265aa", "xn--a1e-rbba",
		"xn--ppl-mra07k", "xn--ple-gdb442a", "xn--apl-izc6365a",
		"xn--ppl-lla2b", "xn--ppl-lla7b", "xn--ple-fsb45p",
		"xn--al-p9b0ca", "xn--a1e-0eda", "xn--app-hvb95v",
		"xn--ppl-svd7624a", "xn--ppl-ura07m", "xn--ple-qbb34e",
		"xn--apl-qoa0m", "xn--apl-qoa0l", "xn--apl-qoa0k",
		"xn--apl-qoa0j", "xn--apl-qoa0i", "xn--apl-t5c00h",
		"xn--ppl-5cd22e", "xn--ppl-pla240c", "xn--apl-jzc59c",
		"xn--al-3kaa0j", "a-pple", "ap-ple",
		"app-le", "appl-e", "aplple",
		"ampple", "apmple", "a0pple",
		"alpple", "appmle", "applke",
		"apople", "apploe", "applme",
		"appole", "appple", "appkle",
		"applle", "ap0ple", "aopple",
		"applpe", "app0le", "appe",
		"appl", "pple", "aple",
		"aapple", "spple", "apmle",
		"applr", "appls", "applw",
		"aplle", "applz", "2pple",
		"apppe", "ample", "ypple",
		"aople", "alple", "appl3",
		"appl4", "1pple", "zpple",
		"wpple", "apole", "appke",
		"appoe", "paple", "aplpe",
		"appel", "appli", "opple",
		"applo", "upple", "applecom",
	}

	return &Brand{
		Name:       name,
		Original:   original,
		Safelist:   safelist,
		Suspicious: suspicious,
	}
}

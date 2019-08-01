// PhishDetect
// Copyright (c) 2018-2019 Claudio Guarnieri.
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

// Netflix brand properties.
func Netflix() *Brand {
	name := "netflix"
	original := []string{"netflix"}
	safelist := []string{
		"netflix.adult", "netflix.af", "netflix.ag", "netflix.ai",
		"netflix.asia", "netflix.at", "netflix-australia.com", "netflix.ax",
		"netflix.berlin", "netflix.bg", "netflix.bi", "netflix.buzz",
		"netflix.bz", "netflix.cat", "netflix.cc", "netflix.ceo", "netflix.cf",
		"netflix.club", "netflix.cm", "netflix.cn", "netflix.co",
		"netflix.co.ag", "netflix.co.at", "netflix.co.bi", "netflix.co.bw",
		"netflix.co.cm", "netflix.co.gl", "netflix.co.gy", "netflix.co.in",
		"netflix.co.ke", "netflix.co.kr", "netflix.co.lc", "netflix.com",
		"netflix.com.af", "netflix.com.ag", "netflix.com.ai", "netflix.com.bi",
		"netflix.com.bz", "netflix.com.cm", "netflix.com.cn", "netflix.com.co",
		"netflix.com.de", "netflix.com.ec", "netflix.com.gl", "netflix.com.gp",
		"netflix.com.gt", "netflix.com.gy", "netflix.com.hk", "netflix.com.hn",
		"netflix.com.ht", "netflix-com.ie", "netflix.com.lc", "netflix.com.lv",
		"netflix.com.ly", "netflix.com.mg", "netflix.com.ms", "netflix.com.mw",
		"netflix.com.ng", "netflix.com.pa", "netflix.com.pe", "netflix.com.pr",
		"netflix.com.ps", "netflix.com.sb", "netflix.com.sc", "netflix.com.sg",
		"netflix.com.sl", "netflix.com.sn", "netflix.com.tj", "netflix.com.tw",
		"netflix.com.ua", "netflix-com.uk", "netflix.com.vc", "netflix.com.vi",
		"netflix.co.mw", "netflix.co.nz", "netflix.co.pn", "netflix.co.tj",
		"netflix.co.tz", "netflix.co.ug", "netflix.co.vi", "netflix.co.za",
		"netflix-customerservice.com", "netflix.cx", "netflix.cz", "netflix.de",
		"netflix.ec", "netflix.film", "netflix.firm.in", "netflix.fm",
		"netflix-forever.de", "netflix.fr", "netflix.ga", "netflix.gd",
		"netflix.gen.in", "netflix.gl", "netflix.gp", "netflix.gq",
		"netflix.gs", "netflix.gy", "netflix.hk", "netflix.hn", "netflix.horse",
		"netflix.ht", "netflix.id", "netflix.in", "netflix-inc.com",
		"netflix.ind.in", "netflix.info", "netflix.jobs", "netflix.jp",
		"netflix.ki", "netflix.kn", "netflix.kr", "netflix.kz", "netflix.la",
		"netflix.lc", "netflix.lu", "netflix.ly", "netflix.md", "netflix.mg",
		"netflix.ml", "netflix.mn", "netflix-service.com",
		"netflix-theater.com", "netflix-theatre.com", "netflix.net",
		"netflix.org",
	}
	suspicious := []string{
		"netflixa", "netflixb",
		"netflixc", "netflixd", "netflixe",
		"netflixf", "netflixg", "netflixh",
		"netflixi", "netflixj", "netflixk",
		"netflixl", "netflixm", "netflixn",
		"netflixo", "netflixp", "netflixq",
		"netflixr", "netflixs", "netflixt",
		"netflixu", "netflixv", "netflixw",
		"netflixx", "netflixy", "netflixz",
		"oetflix", "letflix", "jetflix",
		"fetflix", "ndtflix", "ngtflix",
		"natflix", "nmtflix", "nutflix",
		"neuflix", "nevflix", "nepflix",
		"nedflix", "ne4flix", "netglix",
		"netdlix", "netblix", "netnlix",
		"netvlix", "netfmix", "netfnix",
		"netfhix", "netfdix", "netflhx",
		"netflkx", "netflmx", "netflax",
		"netflyx", "netfliy", "netfliz",
		"netflip", "netflih", "netfli8",
		"xn--net1ix-k6e", "xn--ntflx-b2e03310a", "xn--ntlix-6za67f",
		"xn--netfix-6r6v", "xn--ntfix-joc12u", "xn--ntf1ix-b5a",
		"xn--ntflx-7nc282b", "xn--ntf1ix-33a", "xn--nflix-iza918c",
		"xn--etflix-25c", "xn--neflx-zwb394b", "xn--nelix-nde36l",
		"xn--ntfllx-w4a", "xn--netlx-13a829b", "xn--nefli-8ye0228b",
		"xn--netf1x-6db", "xn--ntfix-uza01v", "xn--netfl-8nc612b",
		"xn--netfl-uze13410a", "xn--netf1i-gfg", "xn--netfx-2sa80d",
		"xn--ntfllx-ph8b", "xn--tflix-v7a733d", "xn--ntfl1x-bva",
		"xn--netl1x-k6e", "xn--nflix-uza508c", "xn--ntflx-iza5934a",
		"xn--netf1x-m91a", "xn--rtflix-b5a", "xn--ntfli-vx1bsn",
		"xn--netllx-y7f", "xn--retflx-m91a", "xn--netfix-6va",
		"xn--ntflx-bsa39b", "xn--nflix-j0a598b", "xn--ntflx-q4a137c",
		"xn--netfx-13aa", "xn--ntfiix-p3a", "xn--neflx-fta860d",
		"xn--eflix-v7a2t", "xn--ntflx-2sa6x", "xn--netfx-n7aa",
		"xn--ntfl1x-iva", "xn--neflx-zwb116a", "xn--neflx-q4a606c",
		"xn--ntflx-2sa702d", "xn--ntfix-v0a86u", "xn--mtflix-b5a",
		"xn--netlx-cmb07260a", "xn--netlx-7nc099a", "xn--retlix-5tb",
		"xn--netlx-fta636d", "xn--ntf1ix-iva", "xn--ntflx-nsa14470a",
		"xn--netli-cmb5303c", "xn--netfx-n7a97r", "xn--ntfllx-p3a",
		"xn--etflix-255b", "xn--ntfli-hsa3124c", "xn--retflx-71c",
		"xn--ntflx-uza79u", "xn--metflx-fze", "netf1ix",
		"xn--nefli-nde2019b", "xn--nflix-7db3863c", "xn--nefllx-qkb",
		"xn--ntf1ix-bhg", "xn--netfl-pjf16600a", "xn--tflix-v7a206c",
		"xn--ntflx-nsa19x", "xn--ntlix-bmb3533c", "xn--netflx-7db",
		"xn--netfl-8nc644a", "xn--ntflx-uza74170a", "xn--ntfl1x-3of",
		"xn--ntlix-iza40g", "xn--netfx-2sa78x", "xn--retlix-k6e",
		"xn--ntfiix-3of", "xn--etflix-205b", "xn--nflix-7ye8358b",
		"xn--netfix-7va", "xn--ntfl1x-p3a", "xn--netlx-5df121y",
		"xn--ntflx-v0a1n", "xn--netlx-cmb279a", "xn--neflx-2sa73f",
		"xn--metfix-6db", "xn--etlix-v7a177b", "xn--neflx-tbe8e",
		"xn--ntfllx-3of", "xn--ntflx-2sa490c", "netfl1x",
		"xn--rtflix-3of", "xn--etflx-v7a96r", "xn--netfl-3sa8814c",
		"xn--ntflx-hsa84470a", "xn--etflix-hqf", "xn--nefllx-qrf",
		"xn--ntflx-zwe187y", "xn--netfx-13a93t", "xn--netfl-23a856c",
		"xn--neflx-8db4383a", "xn--retflx-t9a", "xn--ntflx-fta81a",
		"xn--netli-cmb327c", "xn--metlix-5tb", "xn--metflx-7r6v",
		"xn--netfl-ube49g", "xn--ntflx-bsa5a", "xn--nflix-6za909b",
		"xn--nflix-hsa601d", "xn--ntflx-j0a96u", "xn--tflix-gsa34d",
		"xn--netfx-ftaa", "xn--netlx-2sa673c", "xn--neflx-tbe07g",
		"xn--netflx-m2c", "xn--nflix-19d10h", "xn--netfl-vx1b7886h",
		"xn--mtflix-bva", "xn--ntflx-q4a0304c", "xn--ntfix-joc3381c",
		"xn--ntfiix-iva", "xn--nflix-6za64c", "xn--ntflx-7nc6481c",
		"netfiix", "xn--nelix-nde9o", "xn--ntflx-zwb505a",
		"xn--ntflx-2sa249d", "xn--ntlix-nsa766d", "xn--ntfix-iza54a",
		"xn--ntfli-b2e38d", "xn--netfll-n77b", "xn--netlx-q4a19d",
		"netfllx", "xn--netfii-n77b", "xn--nef1ix-qkb",
		"xn--ntfl1x-pva", "xn--ntlix-xhe4429b", "xn--eflix-v7a774c",
		"xn--mtflix-3of", "xn--netlx-yhe16120a", "xn--metflx-m91a",
		"xn--ntlix-6za580c", "xn--netf1x-08a", "xn--ntflx-6za33170a",
		"xn--netli-cmb359b", "xn--netfl-ube46o", "xn--etflx-q4a2k",
		"xn--ntflx-13a8504c", "xn--ntfli-v0a0783c", "xn--rtflix-iye",
		"xn--tflix-v7a984b", "xn--netfx-n7a92860a", "xn--ntflx-j0a5o",
		"xn--metlix-k6e", "xn--ntfix-m7a316c", "xn--netfl-23a824d",
		"xn--ntflx-v0a48j", "xn--metflx-mwa", "xn--nflix-bsa311d",
		"xn--ntfli-6za8983c", "xn--ntfllx-b5a", "xn--ntflx-v0a9334a",
		"xn--ntflx-v0a3k", "xn--ntflx-uza5o", "xn--rtflix-i4a",
		"xn--ntfli-b2e5908b", "xn--netlx-cmb4053a", "xn--ntfli-j0a4883c",
		"xn--ntflx-tbe4949b", "xn--ntflx-19d5a", "xn--ntflx-iza04k",
		"xn--neflx-2sa002c", "xn--ntfix-6za69u", "xn--retfli-n77b",
		"xn--nflix-7db911d", "xn--nflix-v0a81c", "xn--ntfix-hsa23d",
		"xn--ntfli-hsa109d", "xn--tflix-5za82a", "xn--ntfix-m7a693c",
		"xn--ntflx-j0a7l", "xn--ntflx-j0a91170a", "xn--eflix-v7a595b",
		"xn--nefiix-j1e", "xn--nefli-8db382c", "xn--ntflx-2sa0z",
		"xn--ntflx-iza16170a", "xn--ntfix-bsa93d", "xn--netfl-r4a036c",
		"xn--ntflx-uza3r", "xn--netli-5df0u", "xn--metlix-y7f",
		"xn--ntflx-fta679c", "xn--ntflx-6za21k", "xn--ntfix-nsa52d",
		"xn--netlx-13a202d", "xn--netfx-7ncn", "xn--tflix-v7a1093c",
		"xn--ntfiix-ph8b", "xn--ntlix-zwe37d", "xn--netfl-3sa601d",
		"xn--netfl-0wb782c", "xn--netfx-koc25u", "xn--ntfli-zwe0j",
		"xn--ntfix-joc435a", "xn--netfix-7r6v", "xn--ntlix-19d5z",
		"nnetflix", "xn--netf1i-gsf", "xn--neflx-8db203b",
		"xn--retflx-08a", "xn--ntflx-n51bn532h", "xn--ntlix-nsa48i",
		"xn--ntflx-13a245c", "xn--ntflx-v0a708b", "xn--ntf1ix-iye",
		"xn--netfl-3sa678d", "xn--netfl-0wb715b", "xn--netfix-m91a",
		"xn--ntflx-bsa50y", "xn--ntfli-zwe00g", "xn--tflix-msa63d",
		"xn--ntflx-bsa12c", "xn--netf1x-t9a", "xn--ntlix-nsa393c",
		"xn--ntlix-bmb450c", "xn--ntflx-iza9p", "xn--metflx-m6b",
		"xn--netli-uze75c", "xn--etfli-v7a4853c", "xn--ntfli-uza2193c",
		"xn--netfx-koc4891a", "xn--retflx-7r6v", "xn--retflx-fze",
		"xn--netfl-r4a004d", "xn--nflix-uza329b", "xn--mtflix-i4a",
		"xn--mtflix-bvf", "xn--nflix-iza739b", "xn--ntflx-fta2944c",
		"xn--netfli-uyd75w", "xn--ntflx-uza62k", "xn--nflix-iza47c",
		"xn--ntf1ix-ph8b", "xn--ntflx-hsa2b", "xn--tflix-tza24a",
		"xn--ntflx-q4a664d", "xn--rtflix-bvf", "retfllx",
		"xn--ntflx-k9x250h", "xn--ntfli-hsa131d", "xn--ntfiix-bvf",
		"xn--etflix-o6c", "xn--ntlix-uza09f", "xn--ntfix-joc962b",
		"xn--metflx-08a", "xn--netlx-2sa76i", "xn--ntlix-hsa004c",
		"xn--netfix-z8a", "xn--ntflx-hsa72n", "xn--ntfllx-iva",
		"xn--etflx-v7a125b", "xn--nflix-nsa45f", "xn--reflix-qrf",
		"xn--ntfli-iza416d", "xn--nelix-8db34a", "xn--ntfix-nsa40y",
		"netf1lx", "xn--ntflx-zwb104b", "xn--ntfix-uza13a",
		"xn--ntflx-13a957c", "xn--rtflix-bhg", "xn--ntflx-nsa70c",
		"xn--netlx-fta263c", "xn--netfx-n41sa", "xn--nefix-m7a3u",
		"xn--rtflix-33a", "xn--ntfli-blf0637b", "xn--ntfix-6za71a",
		"xn--ntflx-nsar", "xn--ntlix-bmb738b", "xn--ntfix-m7a005b",
		"xn--ntflx-j0a128b", "xn--tflix-9ra05d", "xn--ntfix-hsa11y",
		"xn--ntfli-j0a208c", "xn--netlx-cmb94b", "xn--netlx-q4a471d",
		"xn--ntflx-bsa9864a", "xn--netfl-pjf598x", "xn--ntlix-v0a133d",
		"xn--ntflx-hsa41c", "xn--retflx-7va", "xn--ntflx-19d34420a",
		"xn--ntflx-uza1834a", "xn--netfix-fze", "xn--netlx-cmb02m",
		"xn--mtflix-w4a", "xn--nef1ix-j1e", "xn--nelix-xhe02e",
		"xn--ntfiix-b5a", "xn--netfix-t9a", "xn--ntfllx-pva",
		"xn--ntflx-7nc43u", "xn--ntfli-j0a275d", "xn--ntflx-nsa97b",
		"xn--nefl1x-j1e", "xn--etflx-2sa81d", "xn--metfli-n77b",
		"xn--ntfix-iza42v", "xn--nflix-6za197c", "xn--mtflix-ph8b",
		"xn--ntlix-b2e65b", "xn--etlix-v7a26c", "xn--netlx-zwb787a",
		"xn--ntfli-pjf3477b", "xn--retflx-mwa", "xn--ntflx-tbe59h",
		"xn--ntfli-19d8029b", "xn--ntflx-6za38u", "xn--meflix-qrf",
		"xn--netix-m7a650d", "xn--net1ix-5tb", "xn--ntflx-6za538b",
		"xn--netfx-zwb89h", "xn--ntfli-uze3158b", "xn--ntfli-iza448c",
		"xn--ntflx-q4a415c", "metfl1x", "xn--neflx-q4a16a",
		"xn--ntfli-6za618c", "xn--rtflix-ph8b", "xn--tflix-u0a00a",
		"xn--ntflx-fta2w", "xn--netfl-gta4714c", "xn--ntlix-uza373d",
		"xn--netfix-s9a", "xn--ntflx-iza368b", "xn--netfx-q4a11t",
		"metf1ix", "xn--ntflx-zwb353c", "xn--ntflx-blf59500a",
		"xn--ntflx-zwb7192c", "xn--nflix-j0a777c", "xn--ntflx-7nc745a",
		"xn--netlx-2sa056d", "xn--etfix-v7a28r", "xn--ntfli-v0a887c",
		"xn--nflix-hsa16f", "xn--netf1i-n77b", "xn--nflix-zwe5g",
		"xn--netfix-08a", "xn--ntflx-blf928x", "xn--nelix-8db628c",
		"xn--ntlix-4df60a", "xn--ntflx-2sa40a", "xn--mtflix-iye",
		"xn--ntlix-xhe82d", "xn--nflix-mde0249b", "xn--netix-cmb33m",
		"xn--metflx-71c", "xn--netfli-uyd735a", "xn--nelix-8db255b",
		"xn--ntfix-v0a9y", "xn--ntflx-fta40a", "xn--nefix-joc904a",
		"xn--etflix-v5c", "xn--reflix-qkb", "xn--retfix-6db",
		"xn--netfi-koc334a", "xn--nefli-8yey", "xn--etfl1x-heb",
		"xn--netfix-mwa", "xn--netf1x-7va", "xn--netflx-6db",
		"xn--ntfl1x-i4a", "xn--etflx-v7a3504a", "xn--neflx-nde73320a",
		"xn--mtflix-bhg", "xn--meflix-qkb", "xn--etflix-oof",
		"xn--ntfllx-bva", "xn--netfi-koc302b", "xn--neflx-7nc04v",
		"xn--netllx-k6e", "xn--nefli-8ye80f", "xn--neflx-13a436c",
		"xn--ntlix-v0a750c", "xn--netflx-l2c", "xn--ntflx-b2e465y",
		"xn--ntf1ix-bvf", "xn--etlix-v7a540d", "xn--netfx-tbea",
		"xn--nflix-7ye47f", "xn--netfix-71c", "xn--netix-m7a37c",
		"xn--net1ix-y7f", "xn--ntfl1x-w4a", "xn--nflix-bsa86f",
		"xn--ntflx-hsa89x", "xn--ntfli-v0a855d", "xn--metfix-l2c",
		"xn--ntf1ix-pva", "xn--etfllx-heb", "xn--netfx-n7a3604a",
		"xn--ntflx-hsa051c", "xn--neflx-nde175z", "xn--rtflix-iva",
		"xn--ntlix-j0a26f", "xn--netf1x-7r6v", "xn--netix-joc30x",
		"xn--ntfiix-bhg", "xn--ntflx-nsa5764a", "xn--nefl1x-qrf",
		"xn--nefix-joc72v", "xn--ntfli-6za685d", "xn--netf1x-fze",
		"xn--nflix-7db172b", "xn--netli-yhe41m", "xn--netfx-q4a2j",
		"xn--ntfli-uze4h", "xn--ntfli-uze94f", "xn--nflix-7ye9j",
		"xn--ntflx-nsa341c", "xn--netiix-5tb", "xn--ntlix-bmb987c",
		"xn--ntflx-hsa2864a", "xn--tflix-i0a41a", "xn--ntfl1x-33a",
		"xn--netlx-q4a009b", "xn--ntfllx-bvf", "xn--netix-joc779a",
		"xn--ntfiix-33a", "xn--nefli-8db350d", "xn--ntfllx-iye",
		"xn--neflx-8db05p", "xn--netfii-gsf", "xn--netfll-gfg",
		"xn--netfl-uze566y", "xn--nefl1x-qkb", "xn--netfl-23a0473c",
		"xn--ntfli-19d69o", "xn--netfx-fta37x", "xn--ntfl1x-bvf",
		"xn--netfx-2saa", "xn--nflix-7db483c", "xn--netfix-lwa",
		"xn--nflix-v0a188b", "xn--netfi-n7a4953c", "xn--ntfllx-i4a",
		"xn--ntflx-6za1n", "xn--meflix-j1e", "retfl1x",
		"rretflix", "xn--ntfix-bsa81y", "xn--ntlix-bsa89i",
		"xn--ntflx-fta080c", "xn--netlx-5df78800a", "xn--rtflix-bva",
		"xn--netfx-zwba", "retf1ix", "xn--netfll-gsf",
		"netf11x", "xn--ntflx-j0a89j", "xn--ntflx-2sa23a",
		"xn--ntfli-19d62h", "xn--netli-5df2567b", "xn--retfix-l2c",
		"xn--ntlix-hsa19i", "xn--netfl-8nc8251c", "xn--retflx-m6b",
		"xn--netfix-l91a", "xn--ntflx-bsa751c", "xn--tflix-hza65a",
		"xn--netfli-uyd5647c", "xn--netfii-gfg", "xn--ntfllx-bhg",
		"xn--ntfix-j0a30a", "xn--ntflx-13a494d", "xn--netfl1-gfg",
		"xn--ntlix-j0a170c", "xn--ntflx-bsa9b", "xn--ntflx-2sa099c",
		"xn--ntf1ix-3of", "xn--ntfli-nsa6024c", "xn--netfi-koc5151c",
		"xn--netli-yhe44e", "xn--netfix-l6b", "xn--neflx-2sa280d",
		"xn--ntflx-19d776z", "xn--ntflx-iza11v", "xn--netix-m7a287b",
		"xn--ntflx-iza7s", "xn--ntfl1x-b5a", "xn--ntflx-fta391d",
		"xn--neflx-8ye55410a", "xn--ntfli-zwe2138b", "xn--netf1x-l2c",
		"xn--ntfli-nsa498d", "xn--ntflx-6za7634a", "xn--netlx-yhe593z",
		"xn--netfl1-n77b", "xn--ntlix-bmb149a", "xn--ntflx-v0a55u",
		"xn--netfl-gta268d", "xn--netiix-y7f", "xn--etf1ix-heb",
		"xn--nflix-uza06c", "xn--nefix-m7a884c", "xn--ntf1ix-w4a",
		"xn--ntfiix-pva", "xn--netli-yhe6298b", "xn--ntflx-tbe87f",
		"xn--nefllx-j1e", "xn--neflx-13a257b", "xn--ntfix-j0a28u",
		"xn--metflx-t9a", "xn--neflx-7nc224a", "xn--nefli-nde09n",
		"xn--nefix-m7a606b", "xn--etflx-v7a91860a", "xn--ntlix-xhe08m",
		"xn--ntflx-fta0z", "xn--ntfix-m7a843d", "xn--netfi-n7a282d",
		"xn--ntfllx-33a", "xn--ntflx-6za9p", "xn--ntfiix-bva",
		"xn--ntfl1x-ph8b", "xn--nefli-8db5633c", "xn--netlx-zwb160c",
		"xn--nelix-8ye18c", "xn--metfli-gsf", "xn--netfx-13a0m",
		"xn--neflx-fta32f", "xn--netfx-fta49c", "xn--ntfli-uza006d",
		"xn--mtflix-p3a", "xn--ntlix-bsa186d", "xn--nflix-hsa422c",
		"xn--metfli-gfg", "xn--neflx-fta681c", "xn--ntflx-zwb816b",
		"xn--ntfli-uza038c", "xn--netfi-n7a215c", "xn--netfix-m6b",
		"xn--netl1x-5tb", "xn--netfx-k9xa", "metfllx",
		"retfiix", "xn--ntflx-uza948b", "xn--neflx-q4a427b",
		"xn--nelix-bmb740b", "xn--ntflx-fta6x", "xn--nefli-nde02g",
		"xn--ntlix-hsa476d", "xn--nflix-v0a367c", "xn--nflix-nsa712c",
		"xn--nefiix-qkb", "xn--neflx-13a98a", "xn--ntfli-pjf9d",
		"xn--retfli-gfg", "xn--netfx-koca", "xn--netfix-61c",
		"xn--ntlix-v0a84f", "xn--mtflix-33a", "xn--ntlix-uza990c",
		"xn--etflx-v7a89g", "xn--nflix-7db761c", "xn--ntflx-nsa5a",
		"xn--etflx-13a0n", "xn--ntfiix-iye", "xn--netfl-gta290d",
		"xn--netfix-eze", "netfi1x", "xn--ntfli-nsa421d",
		"xn--mtflix-iva", "xn--mtflix-pva", "xn--ntflx-v0a50170a",
		"xn--ntflx-fta829d", "xn--netf1x-71c", "xn--etfix-m7ak",
		"xn--nflix-j0a23c", "xn--etflix-o3c", "xn--ntlix-6za953d",
		"xn--neflx-8db97e", "xn--ntfli-bsa809d", "xn--netf1x-m6b",
		"xn--netiix-k6e", "xn--netfl-l9x498g", "xn--netfl-r4a2173c",
		"metfiix", "xn--ntlix-19d96m", "xn--tflix-v7a583c",
		"xn--ntfix-m7a2193c", "xn--ntflx-nsa02n", "xn--etflx-fta40d",
		"xn--etfli-v7a205c", "xn--ntfli-iza6293c", "xn--rtflix-w4a",
		"xn--etfiix-heb", "xn--netfl1-gsf", "xn--ntfiix-i4a",
		"rnetflix", "xn--neflx-8ye986y", "xn--ntfiix-w4a",
		"xn--ntflx-zwe74510a", "xn--nflix-mde65o", "xn--netfx-n7a80h",
		"xn--ntfli-bsa0224c", "xn--ntflx-13a646b", "xn--netllx-5tb",
		"xn--retfli-gsf", "xn--netlx-13a91e", "xn--ntflx-hsay",
		"xn--netfl-ube6719b", "xn--nflix-19d9h", "xn--nefiix-qrf",
		"xn--nelix-bmb929b", "xn--nefix-8db36p", "xn--etflix-2s7b",
		"xn--nflix-nsa990d", "xn--neflx-8db00660a", "xn--reflix-j1e",
		"xn--ntflx-j0a3534a", "xn--ntf1ix-i4a", "xn--rtflix-p3a",
		"xn--retlix-y7f", "xn--netfx-7nca", "xn--netlx-fta35i",
		"xn--netl1x-y7f", "xn--netfx-koc05740a", "xn--ntlix-j0a543d",
		"xn--nflix-mde12h", "xn--ntflx-hsa68b", "xn--ntflx-q4a816b",
		"xn--rtflix-pva", "xn--ntf1ix-p3a", "xn--netfx-q4aa",
		"xn--nflix-mde40f", "netfilx", "xn--ntlix-iza311c",
		"xn--ntflx-bsa43n", "xn--ntfl1x-iye", "xn--netfx-n7a135b",
		"xn--ntfix-joc713a", "xn--netfl-0wb9952c", "xn--ntlix-4df0797b",
		"xn--ntflx-bsa55470a", "xn--ntf1ix-bva", "xn--ntflx-tbe03p",
		"xn--metflx-7va", "xn--ntlix-bsa704c", "xn--ntlix-xhe54f",
		"xn--netlx-tbe4w", "xn--etfli-v7a272d", "xn--ntflx-2sa81a",
		"xn--ntlix-iza783d", "xn--ntfl1x-bhg", "xn--netlx-7nc61x",
		"xn--nflix-bsa132c", "xn--nef1ix-qrf", "xn--netf1x-mwa",
		"xn--ntflx-7nc033a", "xn--netlx-tbe83m", "xn--ntfli-bsa831d",
		"xn--ntflx-2sa6054c", "n-etflix", "ne-tflix",
		"net-flix", "netf-lix", "netfl-ix",
		"netfli-x", "netzflix", "netyflix",
		"ne4tflix", "netflmix", "netvflix",
		"netfmlix", "newtflix", "netfli8x",
		"neytflix", "n3etflix", "netfl9ix",
		"netfl8ix", "nertflix", "nedtflix",
		"netfflix", "negtflix", "netfdlix",
		"nestflix", "ne5tflix", "netflijx",
		"ndetflix", "netfvlix", "netfolix",
		"netcflix", "netftlix", "ne6tflix",
		"netfli9x", "netfglix", "netflpix",
		"neztflix", "neftflix", "netfplix",
		"netrflix", "nettflix", "net6flix",
		"netfluix", "nsetflix", "netflkix",
		"nzetflix", "nwetflix", "netdflix",
		"nretflix", "ne3tflix", "netfrlix",
		"netfljix", "netflikx", "netfclix",
		"netfklix", "net5flix", "netgflix",
		"netfliux", "n4etflix", "netfloix",
		"netfliox", "etflix", "netlix",
		"netfli", "netfix", "neflix",
		"ntflix", "netflx", "netfliix",
		"neetflix", "netfllix", "n3tflix",
		"ne6flix", "ne5flix", "nettlix",
		"nefflix", "netfljx", "netclix",
		"nwtflix", "netrlix", "nezflix",
		"negflix", "netflis", "netfl8x",
		"netfoix", "netfpix", "hetflix",
		"nztflix", "netflux", "nstflix",
		"metflix", "netfl9x", "netflox",
		"netfkix", "betflix", "nerflix",
		"netflic", "nrtflix", "n4tflix",
		"neyflix", "netflid", "entflix",
		"nteflix", "neftlix", "netlfix",
		"netflxi", "nitflix", "notflix",
		"netflex", "netflixcom",
	}

	return &Brand{
		Name:       name,
		Original:   original,
		Safelist:   safelist,
		Suspicious: suspicious,
	}
}

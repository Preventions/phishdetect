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

// Yahoo brand properties.
func Yahoo() *Brand {
	name := "yahoo"
	original := []string{"yahoo"}
	safelist := []string{
		"yahoo.com", "ymail.com", "rocketmail.com", "yahoo.co.uk", "yahoo.fr",
		"yahoo.com.br", "yahoo.co.in", "yahoo.ca", "yahoo.com.ar", "yahoo.com.cn",
		"yahoo.com.mx", "yahoo.co.kr", "yahoo.co.nz", "yahoo.com.hk",
		"yahoo.com.sg", "yahoo.es", "yahoo.gr", "yahoo.de", "yahoo.com.ph",
		"yahoo.com.tw", "yahoo.dk", "yahoo.ie", "yahoo.it", "yahoo.se",
		"yahoo.com.au", "yahoo.co.id", "yahoo.cl", "yahoo.co.jp", "yahoo.co.th",
		"yahoo.com.co", "yahoo.com.my", "yahoo.com.tr", "yahoo.com.ve",
		"yahoo.com.vn", "yahoo.in", "yahoo.nl", "yahoo.no", "yahoo.pl",
		"yahoo.ro", "yahoo.net",
	}
	suspicious := []string{
		"yahooa", "yahoob",
		"yahooc", "yahood", "yahooe",
		"yahoof", "yahoog", "yahooh",
		"yahooi", "yahooj", "yahook",
		"yahool", "yahoom", "yahoon",
		"yahooo", "yahoop", "yahooq",
		"yahoor", "yahoos", "yahoot",
		"yahoou", "yahoov", "yahoow",
		"yahoox", "yahooy", "yahooz",
		"xahoo", "qahoo", "iahoo",
		"9ahoo", "ychoo", "yehoo",
		"yihoo", "yqhoo", "yaioo",
		"yajoo", "yaloo", "yaxoo",
		"yahno", "yahmo", "yahko",
		"yahgo", "yahon", "yahom",
		"yahok", "yahog", "xn--yh-x2a7080aa",
		"xn--ah-zka04za", "xn--aho-izb6727a", "xn--yh-5eb78qa",
		"xn--hoo-lla379b", "xn--yh-pia740aa", "xn--yah-ina496a",
		"xn--yho-qla07n", "xn--yah0-3nd", "xn--yihoo-mra",
		"xn--yh-mia150aa", "xn--yho-gsb5757a", "xn--yho-ina1019a",
		"xn--yao-5xd721q", "xn--yho-una6809a", "xn--yho-hdb632a",
		"xn--yh-jia550aa", "xn--yaiho-581b", "xn--yho-hna70i",
		"xn--yho-hna2019a", "xn--yah-ueda", "xn--yho-6cd12o",
		"xn--yalho-msf", "xn--yho-9oa588c", "xn--yaiho-nye",
		"xn--yao-e7a280b", "xn--yao-see9h", "xn--yho-e7a102b",
		"xn--yao-lob373a", "xn--yah0-3nd", "xn--yh-mia150aa",
		"xn--ya-5wc018xa", "xn--ah-6jaa036a", "xn--yalho-3ta",
		"xn--yho-hdb632a", "xn--hoo-ord4p", "xn--ya-6jaa223b",
		"xn--aho-e7a084a", "xn--yaih-8qaa", "xn--ya-3yaa713u",
		"xn--hoo-8oa355a", "xn--yho-ila514b", "xn--yah-ued95f",
		"xn--ah-gkb4400aa", "xn--yh-jla990ba", "xn--yh-pia740aa",
		"xn--yaih-tqaa", "xn--hoo-koa41o", "xn--aho-hna1c",
		"xn--yho-9ka5i", "xn--yao-ued24m", "xn--yah-una97m",
		"xn--yah-7ge3843a", "xn--yh-mia9fa", "xn--ah0o-4ra",
		"xn--yho-hna2f", "xn--ah-gkb61va", "xn--aoo-koa509b",
		"xn--yho-tvd50g", "xn--ya00-1oe", "xn--yho-ela124b",
		"xn--yho-ela6249a", "xn--yho-hna25m", "xn--yaih-jx5aa",
		"xn--yao-ued24m", "xn--yho-9ka172c", "xn--aho-cfd6905a",
		"xn--yho-mla6149a", "xn--ya-gkaa739b", "xn--yh-7kc763ya",
		"xn--yah0-ngb", "xn--yah-ued9115a", "xn--ylhoo-4ve",
		"xn--yah-una876a", "xn--yah-e7a390c", "xn--yh-kbb74na",
		"xn--yh-6jaa60i", "xn--yh-6jaa1019a", "xn--aihoo-pva",
		"xn--aho-vxc7185a", "xn--ya-gmca73d", "xn--yah-una976a",
		"xn--yh-via1571aa", "xn--yh0o-loa", "xn--yao-5xd79d",
		"xn--yah-czc84d", "xn--yho-9ka1249a", "xn--yho-9ka1349a",
		"xn--yho-qla504b", "xn--aho-koa74m", "xn--yho-hdb10w",
		"xn--yah-5xd3624a", "xn--yah-ued9015a", "xn--yho-mob1o",
		"xn--yah-ued9115a", "xn--ya-gmca279q", "xn--ya0o-15d",
		"xn--hoo-koa902b", "xn--aho-hna153b", "xn--yh00-53d",
		"xn--yho-nnb6877a", "xn--ylhoo-mra", "xn--yh-kbb74na",
		"xn--yah-mob4577a", "xn--yalh-tqaa", "xn--hoo-lla317a",
		"xn--yho-ula6049a", "xn--yho-ula6939a", "xn--yah-ued85f",
		"xn--aho-ued57c", "xn--ah-2vc0ua", "xn--ah-tmc003ya",
		"xn--yho-qla504b", "xn--yah-una86f", "xn--yah-ina38f",
		"xn--aho-ord7z", "xn--ah-h9b27ca", "xn--yho-hdb532a",
		"xn--yalh-evea", "xn--yalho-nye", "xn--yoo-ula929c",
		"xn--yao-ued279q", "xn--aho-mob09b", "xn--yho-hdb180b",
		"xn--yah-czca", "xn--ah-6jaa0c", "xn--yho0-5q5a",
		"xn--yah-e7a32y", "xn--yaiho-3ta", "xn--yah-ina4339a",
		"xn--yoo-9oa478b", "xn--yho-qla077a", "xn--yho-ela550d",
		"xn--yho-9oa052b", "xn--yaiho-nua", "xn--yho-ila514b",
		"xn--yao-tna839b", "xn--hoo-gdb852a", "xn--yao-rdq944f",
		"xn--yho-qla040d", "xn--ya-x0ca79d", "xn--yaiho-j91b",
		"xn--yoo-gsb47v", "xn--yh-yia158ba", "xn--yao-ued279q",
		"xn--aho-koa63f", "xn--aho-koa646a", "xn--yho-hdb00c",
		"xn--yh-6jaa15m", "xn--yh-jia3ga", "xn--yho-hdb178b",
		"xn--yah-5xd40f", "xn--yah-e7a3548a", "xn--aho0-kfc",
		"xn--yh-jia365aa", "xn--yh0o-53d", "xn--yho-qla6g",
		"xn--ah-3yaa020b", "xn--ya00-15d", "xn--yh-6jaa1f",
		"xn--yh00-5na", "xn--aho-cfd63f", "xn--yalho-bec",
		"xn--yalho-uce", "xn--ya-gkaa248c", "xn--ya-gmca73d",
		"xn--aho-koa746a", "xn--yao-una739b", "xn--yho-gsb64p",
		"xn--yh00-boa", "xn--yho-gsb54p", "xn--yah0-xif",
		"xn--yoo-nnb913b", "xn--hoo-gdb87g", "xn--yho-9ka0349a",
		"xn--yoo-6cd90e", "xn--yao-pedf", "xn--yho-ela024b",
		"xn--ah-2vc27fa", "xn--yahoo-i9x", "xn--hoo-mnb39q",
		"xn--yah0-2nd", "xn--aho-izb63m", "xn--yho-4xd2404a",
		"xn--yao-pedg", "xn--yh0o-0na", "xn--ah-kbca01j",
		"xn--yho-9oa590c", "xn--yh-kbb55sa", "xn--yh-jla51ja",
		"xn--yh-pia1571aa", "xn--ah0o-9ld", "xn--yah-ued8115a",
		"xn--yho-bzc2255a", "xn--yao-see0i", "xn--yho-qla5g",
		"xn--yah0-2nd", "xn--yho-e7a1328a", "xn--aho-koa182b",
		"xn--ya-odd490xa", "xn--aho-tedw", "xn--yh-pia9571aa",
		"xn--yaiho-pkg", "xn--yh-mia3671aa", "xn--yho-ela650d",
		"xn--yoo-9oa977c", "xn--yho-czc67c", "xn--yh0o-boa",
		"xn--yho-hdb00w", "xn--yah-ina3339a", "xn--yao-ted34m",
		"xn--yho-6cd6e", "xn--aoo-ord014q", "xn--ah-zka2561aa",
		"xn--yh-gkaa19h", "xn--aho-izb73m", "xn--yh0o-5na",
		"xn--aho-ted67c", "xn--yao-ped55n", "xn--yho-hna752b",
		"xn--yho-ted75e", "xn--ah00-9ld", "xn--yalho-uce",
		"xn--hoo-tla324b", "xn--aho-una533b", "xn--yoo-mla939c",
		"xn--yho-hna270c", "xn--yh00-0na", "xn--yh-kbb55sa",
		"xn--aho-ord7554a", "xn--yho-6cd6e", "xn--yoo-ela414b",
		"xn--yho-ina1f", "xn--yho-qla077a", "xn--ah-6jaa053b",
		"xn--yoo-9ka459c", "xn--yah-ina396a", "xn--yah-una876a",
		"xn--yho-ued65e", "xn--yoo-ila930c", "xn--yho-ula542c",
		"xn--yah-una313b", "xn--yh-kbb9540aa", "xn--yh-9bba1257a",
		"xn--yho-6cd03g", "xn--yah-czca", "xn--yh-yia530aa",
		"xn--yahoo-0sa", "xn--yao-bzc87s", "xn--yalh-y0ba",
		"xn--yah-ina49m", "xn--yao-ped4215a", "xn--ya00-meu",
		"xn--yaiho-mua", "xn--yh-6yc72ea", "xn--yah0-y0b",
		"xn--aho-cfd6805a", "xn--yao-tna348c", "xn--yh-gkaa63m",
		"xn--yao-7ge733p", "xn--yao-ued73d", "xn--aho-bzc11j",
		"xn--yaih-3nda", "xn--aho0-u6d", "xn--yho-ila1149a",
		"xn--yalho-pkg", "xn--yho-6cd13g", "xn--yh00-qoa",
		"xn--yho-tvd5h", "xn--yao0-1oe", "xn--yh-jia55ka",
		"xn--yho0-gse", "xn--ah-gmca57c", "xn--ya-x0ca721q",
		"xn--hoo-cfd43e", "xn--yho-ela6g", "xn--yao-ped46f",
		"xn--yho-tna29h", "xn--yao-lob872b", "xn--ya0o-meu",
		"xn--yh-kbca19k", "xn--yho-czc19k", "xn--aho-koa646a",
		"xn--yaiho-nye", "xn--hoo-pla824b", "xn--yh-mia15ka",
		"xn--yahoo-xwb", "xn--yoo-mla404b", "xn--yh-yia530aa",
		"xn--yah-ina369c", "xn--alhoo-muc", "xn--yh-jia1fa",
		"xn--yho-9ka524b", "xn--yho-9oa5529a", "xn--yh-via7ea",
		"xn--ah-zka04ja", "yalhoo", "xn--yoo-qdqz13f",
		"xn--yho-bzc2255a", "xn--yho-6cd02o", "xn--aho-vxc65m",
		"xn--yh-6yc5fa", "xn--yah-ina396a", "xn--aho-d7a120b",
		"xn--ya-kbca77s", "xn--yho-9oa152b", "xn--hoo-vxc4855a",
		"xn--yho-qla1049a", "xn--yho-9ka197a", "xn--yho-qla1149a",
		"xn--yho-9oa516a", "xn--yah-e7aa", "xn--yh-6jaa170c",
		"xn--yao-ped56f", "xn--yho-ila1249a", "xn--yao-d7a813u",
		"xn--yao-ted34m", "xn--hoo-8ka86s", "xn--yh-jla137ba",
		"xn--yahoo-l41s", "xn--yihoo-kwb", "xn--yalho-lsf",
		"xn--yho-ula567a", "xn--yh-jia550aa", "xn--aho-czc07d",
		"xn--yah-czc3475a", "xn--yao-3sd4944a", "xn--yalho-581b",
		"xn--yho-sgzvk", "xn--hoo-gdb897a", "xn--yah0-ogb",
		"xn--ya-gkaa271v", "xn--hoo-ord4234a", "xn--yah-ina48f",
		"xn--yao-czc27j", "xn--hoo-koa95h", "xn--hoo-koa40m",
		"xn--yh-sia155aa", "xn--yah-ued85f", "xn--yao-bzc34d",
		"xn--yah-vlzj", "xn--yho-sgz6j", "xn--hoo-vxc93d",
		"xn--yh-7kc97ka", "xn--yh-kbb13za", "xn--yh-kbca1255a",
		"xn--ah-h9b868ya", "xn--hoo-pla869b", "xn--yho-ila150d",
		"xn--hoo-koa4a", "xn--yalho-681b", "xn--yao-bzc87s",
		"xn--yalho-j91b", "xn--yho-czc19k", "xn--yao-rdq044f",
		"xn--yh-3yaa683a", "xn--yao-rdqt54f", "xn--yoo-nnb946t",
		"xn--ah-gkb04oa", "xn--yho-ula104b", "xn--yho-ina16o",
		"xn--ya-wcm266ea", "xn--yho-una64o", "xn--yah-5xd4724a",
		"xn--aho-bzc11j", "xn--ah-3yaa06x", "xn--yh-pia3ea",
		"xn--yah0-7qa", "xn--yh-via93ka", "xn--yho-ula5939a",
		"xn--yah-ueda", "xn--yah0-ix5a", "xn--yho-gsb64p",
		"xn--aho-vxc21e", "xn--yh-jla74da", "xn--yho-nnb19x",
		"xn--yah-mob94x", "xn--yh-5eb593aa", "xn--yho-mla652c",
		"xn--yho-ula104b", "xn--yho-ula567a", "xn--yah-e7a42e",
		"xn--ya-kbca77s", "xn--yho-9ka624b", "xn--ya-9bba273a",
		"xn--yho-ela5149a", "xn--hoo-lla334b", "xn--yah-ueda",
		"xn--yho-qla140d", "xn--yahoo-5nc", "xn--yah-czc39l",
		"xn--aho-vxc21e", "xn--yho-bzc29k", "xn--yho-nnb19x",
		"xn--yho-una142b", "xn--yah-mob94x", "xn--yah-czc39l",
		"xn--aho-ord6454a", "xn--yh-x2a30ra", "xn--yho0-q5b",
		"xn--yho-nnb636a", "xn--yh-x2a30ra", "xn--yalh-tx5aa",
		"xn--aho0-4ra", "xn--yah-inao", "xn--aho-koa73f",
		"xn--yao-tna371v", "xn--yah-czca", "xn--ah-h9b4ca",
		"xn--yah-ueda", "xn--yaloo-ify", "xn--aho-izb17s",
		"xn--yihoo-rwa", "xn--yho-ila0149a", "xn--yho-ula642c",
		"xn--yho-ela6149a", "xn--yho-ela124b", "xn--yoo-oed36e",
		"xn--yah-czc38t", "xn--yho-qla604b", "xn--ah-3yaa084a",
		"xn--yao-see0953a", "xn--yho-mla114b", "xn--yao-see0853a",
		"xn--yh-via930aa", "xn--yh-3yaa18d", "xn--yho-qla0049a",
		"xn--yah-czca", "xn--yah-czc49l", "xn--aho-vxc75m",
		"xn--aho-koa6729a", "xn--yao-ted379q", "xn--yalh-3nda",
		"xn--yho-mla5149a", "xn--ah-zka27da", "xn--yh-x2a11wa",
		"xn--yho-nnb55r", "xn--aoo-izb59y", "xn--yaloo-1gg",
		"xn--yah-czc38t", "xn--yho-9oa516a", "xn--aho-lob19b",
		"xn--yh-7kc1da", "xn--yh-mia5571aa", "xn--yoo-nnb48x",
		"xn--yho-9oa616a", "xn--yho-mob1257a", "xn--ya-3yaa280b",
		"xn--yho-nnb5877a", "xn--yaiho-681b", "xn--yho-ila0g",
		"xn--aho-koa7829a", "xn--aoo-vxc076r", "xn--ya-odd680xa",
		"yallhoo", "xn--yho-qla06g", "xn--hoo-izb410a",
		"xn--yho-hdb1428a", "xn--aho-d7a10j", "xn--yho-hdb1328a",
		"xn--ah-zka657ba", "xn--ah-zka844aa", "xn--aho-czc07d",
		"xn--ya-hfda733p", "xn--yh-kbca1255a", "xn--yh-5eb97la",
		"xn--ah-gkaa5a", "xn--aihoo-dze", "xn--yoo-ula420c",
		"xn--yalh-3nda", "xn--hoo-hla879b", "xn--yoo-9ka482v",
		"xn--ah-zka04za", "xn--ya-5wc0pa", "xn--yho-gsb6757a",
		"xn--yh-5eb36xa", "xn--yao-d7a380b", "xn--ah-3yaa00j",
		"xn--yah-ued94n", "xn--yah-una849c", "xn--yoo-ila472v",
		"xn--ah-h9b27ca", "xn--aoo-oed2a", "xn--aho-vxc7d",
		"xn--yho-9ka160d", "xn--yh-f5s7ia", "xn--yah-mob383b",
		"xn--yaloo-0kf", "xn--yh-3yaa102b", "xn--yho-e7a19f",
		"xn--yho-9oa6429a", "xn--yh0o-qzb", "xn--yho-9oa6529a",
		"xn--hoo-8ka827a", "xn--ya00-wff", "xn--aho-vxc74u",
		"xn--ah00-4ra", "xn--aho-vxc7085a", "xn--yho-ula004b",
		"xn--yho-hdb0428a", "xn--yh00-gse", "xn--yho-nnb624b",
		"xn--aoo-vxc53k", "xn--yh-via930aa", "xn--yho-ula630d",
		"xn--ah-h9b4ca", "xn--yh-via17ea", "xn--yah-mob41r",
		"xn--aho-vxc6085a", "xn--ylhoo-3qa", "xn--yihoo-9qa",
		"xn--yah-mob483b", "xn--yoo-oed3984a", "xn--yho-nnb524b",
		"xn--yh-x2a11wa", "xn--yho-qla604b", "xn--yoo-hdb467b",
		"xn--yah-mob31r", "xn--yho-6cd5e", "xn--yah-inap",
		"xn--yao-hna858c", "xn--yho-9ka097a", "xn--ya0o-wff",
		"xn--aho-e7a00j", "xn--hoo-koa420c", "xn--yao-una703b",
		"xn--yh0o-qoa", "xn--yh00-qzb", "xn--aho-izb17s",
		"xn--hoo-izb4407a", "xn--yho-mla540d", "xn--ya-odd8fa",
		"xn--yh-yia7471aa", "xn--hoo-gdb83v", "xn--yao-ted379q",
		"xn--yihoo-3qa", "xn--yho-ila614b", "xn--yho-nnb55r",
		"xn--yho-czc1255a", "xn--yho-e7a6t", "xn--yah-5xd3724a",
		"xn--yho-9oa052b", "xn--yaih-evea", "xn--yh-jia932ba",
		"xn--ya-kbca706r", "xn--ya-dmc62ka", "xn--yho-d7a28d",
		"xn--yah-czc4475a", "xn--yh-6yc327xa", "xn--yoo-2sd3j",
		"xn--yahoo-o4a", "xn--yho0-goa", "xn--yh-mia38ea",
		"xn--yoo-6cd440r", "xn--aho-hna198b", "xn--yaih-85da",
		"xn--yh-x2a782aa", "xn--yho-ula5e", "xn--yahoo-ioc",
		"xn--yho-ila614b", "xn--yh0o-gse", "xn--yh-5eb78qa",
		"xn--yho-czc1255a", "xn--ah-tmc40ka", "xn--yho-ila162c",
		"xn--yah0-tqa", "yah00", "xn--yho-6cd1815a",
		"xn--ya-9bba772b", "xn--yho-6cd1915a", "xn--aho-ued57c",
		"xn--yho-9oa152b", "xn--aho-d7a16x", "xn--yah-7ge4743a",
		"xn--hoo-mnb353a", "xn--yho-ela58n", "xn--yho-ila062c",
		"xn--yaiho-i91b", "xn--yoo-ula493b", "xn--yah-czc84d",
		"xn--yho-ila087a", "xn--yho-9oa50f", "xn--yoo-gsb402a",
		"xn--yho-ela587a", "xn--yalh-yifa", "xn--yah-inaa",
		"xn--yho-6ge2523a", "xn--ya-dmc413ya", "xn--aoo-izb56s",
		"xn--yho-mob194a", "xn--yh0o-5q5a", "xn--yao-mob273a",
		"xn--yah-vlzk", "xn--yh-gmca6884a", "xn--yao-d7a354a",
		"xn--yaih-85da", "xn--yho-9oa5429a", "xn--yh-hfda1523a",
		"xn--aho-mob07x", "xn--yah-una8139a", "xn--yh-sia7471aa",
		"xn--aho-ina098b", "xn--yho-ela024b", "xn--ah-gkb04oa",
		"xn--yh-gmca65e", "xn--yho-ula004b", "xn--yao-3sd42h",
		"xn--yalho-mye", "xn--yho-tna74o", "xn--yh-mia955aa",
		"xn--hoo-8oa318b", "xn--yoo-9ka940c", "xn--yah-3lza",
		"xn--alhoo-vif", "xn--ah-gmcav", "xn--yah-e7a42y",
		"xn--yaiho-msf", "xn--aho-tna633b", "xn--yalho-9dc",
		"xn--yao-ped5215a", "xn--yao-hna323b", "xn--yao-ped5115a",
		"xn--yho-mla552c", "xn--aoo-koa572b", "xn--ya-6jaa758c",
		"xn--yho-mla014b", "xn--yah-ued94n", "xn--yho-9ka097a",
		"xn--ah-gkaa516a", "xn--yh-yia1da", "xn--yho-mla640d",
		"xn--ah-6jaa098b", "xn--yh-gmca6884a", "xn--ya0o-1oe",
		"xn--yah-e7a42y", "xn--yho-d7a783a", "xn--yh-6jaa652b",
		"xn--yho-9ka09n", "xn--yao-6ge833p", "xn--yho-ila087a",
		"xn--yh-jia365aa", "xn--yah-ued8015a", "xn--yao-bzc806r",
		"xn--yah-mob4477a", "xn--yh-5eb97la", "xn--yh-jla9161aa",
		"xn--yao-una271v", "xn--yho-nnb6977a", "xn--yho-5xd1404a",
		"xn--aho-czc01j", "xn--yao-ina781v", "xn--ah-zka4461aa",
		"xn--yalho-4ta", "xn--yho0-boa", "xn--yh-via3471aa",
		"xn--yalho-nua", "xn--hoo-8ka8l", "xn--aho-izb711a",
		"xn--ah00-9me", "xn--yh-6jaa16o", "xn--yh-sia5571aa",
		"xn--yho-lob294a", "xn--yho0-0mb", "xn--hoo-5cd84d",
		"xn--yao-ted83d", "xn--yoo-ila449c", "xn--aho-czc01j",
		"xn--yah-czc48t", "xn--aho-izb7627a", "xn--ah0o-u6d",
		"xn--aho-izb7727a", "xn--yho-ela1i", "yaiho0",
		"xn--yah-mob31r", "xn--yaih-3nda", "xn--yah-e7a32y",
		"xn--yh-pia368ba", "xn--hoo-tla3j", "xn--aho-lob17x",
		"xn--yoo-ela949c", "xn--yho-hdb00w", "xn--yh-gkaa64o",
		"xn--yho-tna242b", "xn--aho-ord6z", "xn--yho-gsb5857a",
		"xn--yho-mla677a", "xn--yah-una9139a", "xn--ylhoo-0jc",
		"xn--yah-czc48t", "xn--yh00-loa", "xn--yoo-9oa442b",
		"xn--yho-nnb6b", "xn--yoo-qla462v", "xn--yah-moba",
		"xn--yh-pia555aa", "xn--yho-ula66n", "xn--yah-e7a954a",
		"xn--yh-sia57ea", "xn--ylhoo-tcc", "xn--yho-bzc77c",
		"xn--aho-ina0c", "xn--aho-d7a184a", "xn--yao0-wff",
		"xn--hoo-fsb32a", "xn--aho-tna65r", "xn--aho-koa746a",
		"xn--yah0-yif", "xn--yh-sia340aa", "xn--yh-x2a30ba",
		"xn--ya-3yaa789b", "xn--aho-mob05q", "xn--yh-via745aa",
		"xn--yah-5xd4624a", "xn--yao-4xd89d", "xn--yihoo-xqa",
		"yaihoo", "xn--yah-e7a954a", "xn--yh-kbb1540aa",
		"xn--ah00-u6d", "xn--yho-9ka19n", "xn--yao-tna803b",
		"xn--yh-3yaa6t", "xn--yho-9oa60f", "xn--aho-koa7729a",
		"xn--yh-gkaa650c", "xn--yho-gsb6857a", "xn--yah-mob84x",
		"xn--yao-see9753a", "xn--yho-hdb532a", "xn--ah-kbca07d",
		"xn--yh-f5szia", "xn--yho-mla577a", "xn--yho-hna26o",
		"xn--ah-kbca07d", "xn--yh0o-q5b", "xn--yho-ela68n",
		"xn--yh-pia555aa", "xn--aho-vxc11e", "xn--yho-qla177a",
		"xn--yh-5eb3820aa", "xn--yho-ela0i", "xn--yho-mla67n",
		"xn--hoo-5cd8g", "xn--yao-ina758c", "xn--yho-qla177a",
		"xn--yao-una248c", "xn--yah-mob84x", "xn--yihoo-0jc",
		"xn--yah-e7a3648a", "xn--yh-gkaa142b", "xn--yh-9bba67w",
		"xn--ya-5wc818xa", "xn--yah-e7a403b", "xn--yoo-qla920c",
		"xn--yalh-jx5aa", "xn--yaiho-lsf", "xn--ylhoo-9qa",
		"xn--aho-tna6a", "xn--yah-una9039a", "xn--yoo-gsb901b",
		"xn--yao-e7a254a", "yalh0o", "xn--yalho-xob",
		"xn--yalh-85da", "xn--yah0-dve", "xn--yaiho-xob",
		"xn--yh-jla324aa", "xn--yh-jia9571aa", "xn--ya-dmcea",
		"xn--yh00-5q5a", "xn--yao0-meu", "xn--ah-gkaa55r",
		"xn--yh-6yc517xa", "xn--yho-ila08n", "yalho0",
		"xn--yao-bzc806r", "xn--aho-cfd72n", "xn--yho-bzc29k",
		"xn--yahoo-rbe", "xn--yao-3sd52h", "xn--yihoo-2of",
		"xn--yho-9ka0h", "xn--yh00-q5b", "xn--yho-ila187a",
		"xn--yao-ted83d", "xn--aho-ina053b", "xn--yho-tna73m",
		"xn--aho-izb73m", "xn--ya-dmc49da", "xn--yao-rdqz44f",
		"xn--yho-d7a2328a", "xn--yh-7kc573ya", "xn--ah-kbca01j",
		"xn--yho-e7a683a", "xn--yh-yia9371aa", "xn--yoo-gsb934t",
		"xn--aho-ord7454a", "xn--yaiho-4ta", "xn--ya-gkaa703b",
		"xn--aho-vxc11e", "xn--yaih-y0ba", "xn--yho-bzc77c",
		"xn--yho-ela587a", "xn--yho-ila187a", "xn--yho-ued6884a",
		"xn--aoo-koa018c", "xn--aho-lob113a", "xn--yho-nnb09x",
		"xn--yh0o-0mb", "xn--yh-7kc74ea", "xn--yho-mla677a",
		"xn--yho-qla17n", "xn--hoo-dla36s", "xn--yah-czc94d",
		"xn--ylhoo-kwb", "xn--yah-czc84d", "xn--ya-dmc223ya",
		"xn--yho-mla0h", "xn--yho-tvd60g", "xn--yoo-hdb958a",
		"xn--hoo-hla8k", "xn--ylhoo-xqa", "xn--hoo-fsb38o",
		"xn--yah-e7a490c", "xn--yho-ela687a", "xn--aoo-ord08g",
		"xn--aoo-vxc04t", "xn--yihoo-gra", "xn--yah0-sqa",
		"xn--yao0-15d", "xn--yho-ued6884a", "xn--yho-9ka624b",
		"xn--yah-e7a303b", "xn--yah-czc94d", "xn--yh-3yaa1328a",
		"xn--yho-una6d", "xn--yho-una19h", "xn--yho-tna750c",
		"xn--yah0-sx5a", "xn--yao-4xd821q", "xn--aho-vxc6d",
		"xn--yah-unaa", "xn--yho-ela687a", "xn--yah-ued84n",
		"xn--yho-mla6049a", "xn--yho-ina170c", "xn--yah-una8039a",
		"xn--yah-ued95f", "xn--yho-ila18n", "xn--yh-mia955aa",
		"xn--yho-nnb5b", "xn--yah-czc94d", "xn--yaioo-ify",
		"xn--yho-gsb54p", "xn--aihoo-muc", "xn--yah-7ge3743a",
		"xn--hoo-lla3k", "xn--yao-e7a713u", "xn--yho-tvd5634a",
		"xn--ah-h9b84ja", "xn--yalho-yob", "xn--hoo-dla3l",
		"xn--yh00-0mb", "xn--hoo-izb99r", "xn--yao-czc27j",
		"xn--yh-9bba194a", "xn--yao-3sd5t", "xn--yho-nnb09x",
		"xn--aho-ina036a", "xn--aho-lob15q", "xn--yalho-nye",
		"xn--yah-czc94d", "xn--ya-dmcea", "xn--yho-tvd6h",
		"yalh00", "xn--yaih-ogba", "xn--yalh-ogba",
		"xn--yoo-9oa901v", "xn--yoo-hdb922a", "xn--yho0-53d",
		"xn--yoo-qla993b", "xn--yah-ina4239a", "xn--yah0-tx5a",
		"xn--yaih-yifa", "xn--yah0-jx5a", "xn--aho-cfd73f",
		"xn--hoo-dla327a", "xn--aho-una5a", "xn--yho-mla5f",
		"xn--yh-x2a5180aa", "xn--aho-uedv", "xn--yah-7gea",
		"xn--yao-bzc34d", "xn--hoo-hla85s", "xn--yho-mla57n",
		"xn--yahoo-dta", "xn--yah-5xda", "xn--yh-yia53ka",
		"xn--hoo-tla369b", "xn--aoo-izb008a", "xn--yah-czc49l",
		"xn--aho-mob013a", "xn--aho-koa64m", "xn--yao-bzc37j",
		"xn--ya-9bba706t", "xn--ah-9bba013a", "xn--yh-via745aa",
		"xn--yho-9ka072c", "xn--yao-mob772b", "xn--aho-vxc6185a",
		"xn--yho-6cd0915a", "xn--aho-uedv", "xn--ya-gmca24m",
		"xn--yoo-9ka914b", "xn--aihoo-vif", "xn--aho-koa282b",
		"xn--yao-d7a889b", "xn--aho-cfd62n", "xn--yh-x0ca1404a",
		"xn--ylhoo-rwa", "xn--ylhoo-rqa", "xn--yihoo-j11b",
		"xn--aho-ord78h", "xn--ah-gkb6300aa", "xn--yho-qla0f",
		"xn--yalho-i91b", "xn--yho-7ge1523a", "xn--yho-ina652b",
		"xn--yao-ued73d", "xn--aoo-izb031t", "xn--yaiho-yob",
		"xn--yho-ula667a", "xn--aho-ord6554a", "xn--yho-9ka197a",
		"xn--yah-5xd30f", "xn--yah-czc3475a", "xn--ya-kbca27j",
		"xn--yao-e7a789b", "xn--aho-vxc64u", "xn--hoo-mnb33c",
		"xn--yah-czc84d", "xn--yho-ula530d", "xn--yho0-qoa",
		"xn--ylhoo-j11b", "xn--aho-vxc6d", "xn--ya-kbca706r",
		"xn--yho-nnb65r", "xn--yh-9bba1o", "xn--yh-via558ba",
		"xn--yoo-tvd99e", "xn--yho-ula667a", "xn--yah-vlza",
		"xn--yao-ina259b", "xn--yh-sia34ka", "xn--aho-izb6627a",
		"xn--yho0-0sa", "xn--yah-mob3477a", "xn--hoo-dla389b",
		"xn--hoo-8ka889b", "xn--ya-wcmv76ea", "xn--yho-ula5049a",
		"xn--ah0o-9me", "xn--yh-sia1fa", "xn--ah-gkb23ja",
		"xn--aho-vxc7d", "xn--ya-9bba24x", "xn--yihoo-rqa",
		"xn--yho-d7a202b", "xn--yho-nnb65r", "xn--yah-mob41r",
		"xn--hoo-pla8j", "xn--yho-gsb624a", "xn--aho-hna136a",
		"xn--ah-zka844aa", "xn--yh-kbb365aa", "xn--yho-hdb10c",
		"xn--ah-2vc868xa", "xn--yh-yia912ba", "xn--yh-jia78ea",
		"xn--hoo-fsb30w", "xn--aho-tna678b", "xn--yh-yia76ea",
		"xn--yho-9oa51m", "xn--yah-ina496a", "xn--hoo-tla34s",
		"xn--aho-e7a020b", "xn--hoo-koa4509a", "xn--yihoo-tcc",
		"xn--yh-via322ba", "xn--yh-gkaa6d", "xn--aho-izb611a",
		"xn--yao-ped4115a", "xn--yao-ina223b", "xn--aoo-koa041v",
		"xn--yho-qla152c", "xn--aoo-cfd059q", "xn--yho-qla0149a",
		"xn--yho-mob67w", "xn--yho-sgzwk", "xn--yho-sgz7j",
		"xn--yho-ted75e", "xn--yaiho-bec", "xn--yao-pedg",
		"xn--ya-kbca24d", "xn--yho-hdb10w", "xn--yao-bzc37j",
		"xn--aho-izb63m", "xn--yho-mla5049a", "xn--ylhoo-2of",
		"xn--yah-ina823b", "xn--yh-5eb1920aa", "xn--yoo-ree8533a",
		"xn--hoo-hla817a", "xn--yah-una413b", "xn--hoo-pla84s",
		"xn--hoo-dla344b", "xn--ya-kbca27j", "xn--yho-mla6f",
		"xn--yao-pedf", "xn--yoo-6cd41n", "xn--alhoo-dze",
		"xn--hoo-lla35s", "xn--yah0-x0b", "xn--yho0-qzb",
		"xn--yho-ila050d", "xn--yh-jla51za", "xn--yho-ela5g",
		"xn--hoo-fsb341a", "xn--yao-ped45n", "xn--yao-hna881v",
		"xn--yah0-75d", "xn--aoo-oed28c", "xn--yho-ila5h",
		"xn--yho-9ka0249a", "xn--yho0-0na", "xn--yoo-ela440c",
		"xn--aho-bzc17d", "xn--yoo-ula952v", "xn--ah-zka421ba",
		"xn--aho-una55r", "xn--yho-gsb524a", "xn--yh-gmca65e",
		"ya1hoo", "xn--aho0-9ld", "xn--yho-mla577a",
		"xn--yho-ina60i", "xn--yho-hdb078b", "xn--ah-gkb23ja",
		"xn--aho-koa282b", "xn--yah0-75d", "xn--yah-e7a4648a",
		"xn--yho-qla16g", "xn--yah-mob3577a", "xn--yaioo-fye",
		"xn--yho-ued65e", "xn--yho-ula1g", "xn--yho-9oa616a",
		"xn--yho-gsb18v", "xn--yho-ula6e", "xn--yh-sia958ba",
		"xn--aoo-cfd02m", "xn--hoo-vxc45l", "xn--yh-jla324aa",
		"xn--yho-una63m", "xn--yah-una87m", "xn--yho-6cd5e",
		"xn--yho-ina15m", "xn--yh-7kc1da", "xn--yh-mia768ba",
		"xn--yah-czc3575a", "xn--yah-una951c", "xn--yho-9ka524b",
		"xn--alhoo-o9d", "xn--yah-ina823b", "xn--yho-nnb536a",
		"xn--yalho-okg", "xn--yah-mob395a", "xn--yah-una313b",
		"xn--yao-czc24d", "xn--yah-ued8115a", "xn--yho-9ka1h",
		"xn--yh-sia722ba", "yaih0o", "xn--yao-3sd4t",
		"xn--yho-9ka6i", "xn--yaiho-9dc", "xn--ah-gmca57c",
		"xn--yah-ued84n", "xn--yoo-mla962v", "xn--yho-tna7d",
		"xn--yh-sia9da", "xn--hoo-pla807a", "xn--yah-czc3575a",
		"xn--aho-una516a", "xn--yho-ila07g", "xn--ya-6jaa781v",
		"xn--yh-mia7ea", "xn--yho-ila17g", "xn--yah-ued8015a",
		"xn--yho-una650c", "xn--ah-9bba07x", "xn--yh-kbca19k",
		"xn--yho-lob2257a", "xn--yho-9oa688c", "xn--ya-kbca24d",
		"xn--yho-qla052c", "xn--aihoo-o9d", "xn--hoo-hla834b",
		"xn--yh-jla51za", "xn--yho-6cd0815a", "xn--aho-izb27s",
		"xn--yao-mob706t", "xn--yaiho-vce", "xn--yah-ina469c",
		"xn--yah-una976a", "xn--yho-lob77w", "xn--ah-tmc27da",
		"xn--ya-3yaa254a", "xn--ya-gmca24m", "xn--aho-izb609a",
		"xn--yh-3yaa19f", "xn--aho-ted67c", "xn--hoo-mnb31y",
		"yaih00", "xn--aho-ord68h", "xn--yho-czc67c",
		"xn--yh-sia340aa", "xn--yoo-mla430c", "xn--aho-izb27s",
		"xn--yaiho-okg", "xn--yoo-ila904b", "xn--yh-pia5fa",
		"xn--yh00-goa", "xn--yho-9ka060d", "xn--yh-sia155aa",
		"xn--yh-via5da", "xn--yho-ela67g", "xn--ya-gmca279q",
		"xn--yh-kbca67c", "xn--yah-e7a4548a", "xn--aho-ina07r",
		"yaiihoo", "xn--yh-kbb7aa", "xn--yh-pia97ea",
		"xn--hoo-8oa372b", "xn--yao-mob24x", "xn--yah-e7a854a",
		"xn--yah-una96f", "xn--ah-gkb841aa", "xn--yho0-loa",
		"xn--yah-ina371c", "xn--yh-pia132ba", "xn--yho-mla66g",
		"xn--yh-jia178ba", "xn--yoo-ela972v", "xn--yah-mob495a",
		"xn--aho-koa621c", "xn--yaiho-mye", "xn--yh-yia3ea",
		"xn--yho-mla114b", "xn--ah-gmcav", "xn--ah-9bba05q",
		"xn--yho-nnb5977a", "xn--yah-una949c", "xn--alhoo-pva",
		"xn--ah-h9b08pa", "xn--yh-pia74ka", "xn--yh00-0sa",
		"xn--aho-koa721c", "xn--hoo-cfd4684a", "xn--yahoo-l7a",
		"xn--yho-ula65g", "xn--yoo-nnb414a", "xn--yao-hna359b",
		"xn--ah-tmc292ya", "xn--yah-ued9015a", "xn--aho-izb709a",
		"xn--ya-5wc22fa", "xn--yho-tvd6634a", "xn--yalho-vce",
		"xn--yho-tvd6734a", "xn--yao-3sd5844a", "xn--yao-3sd4844a",
		"xn--yao-3sd5944a", "xn--yao-czc706r", "xn--yalho-mua",
		"xn--yho-gsb08v", "xn--aho-una578b", "xn--yh-x2a919aa",
		"xn--yah-ina471c", "xn--aoo-cfd51d", "xn--yh-yia345aa",
		"xn--yah0-eve", "xn--aho-cfd7805a", "xn--yho-tna7809a",
		"xn--aho-cfd7905a", "xn--yaioo-1gg", "xn--hoo-tla307a",
		"xn--yao-czc24d", "xn--yho-mla014b", "xn--ah-gkaa533b",
		"xn--yaloo-fye", "xn--yihoo-4ve", "xn--yho-ela562c",
		"xn--yalho-vce", "xn--yah-ina39m", "xn--aoo-ord5d",
		"xn--yh0o-0sa", "xn--yao-czc706r", "xn--yho0-5na",
		"xn--yho-ela662c", "xn--yah-una851c", "xn--ah-6jaa07r",
		"xn--yalho-mye", "xn--yh-jla7261aa", "xn--ah-gkaa578b",
		"xn--yho-9oa690c", "xn--aho-koa182b", "xn--aho-tedw",
		"xn--yaiho-uce", "xn--yh-kbca67c", "xn--yho-d7a7t",
		"xn--yah-ina923b", "xn--aho0-9me", "xn--yh0o-goa",
		"xn--ah-9bba09b", "xn--yah0-85d", "xn--yh-mia532ba",
		"xn--yho-ila6h", "xn--yoo-hdb490u", "xn--yh-yia345aa",
		"xn--yho-ila1g", "xn--ah0o-kfc", "xn--yoo-tvd922q",
		"xn--yho-mla1h", "xn--aho-hna17r", "xn--yho-lob2o",
		"xn--yho-ted7884a", "xn--yoo-qla439c", "xn--yaiho-vce",
		"xn--ya-6jaa259b", "xn--ylhoo-gra", "xn--ah-h9b678ya",
		"xn--yah-ina923b", "xn--yaioo-0kf", "xn--yoo-2sd3624a",
		"xn--yah0-85d", "xn--yahoo-z3a", "xn--hoo-8oa39q",
		"xn--yho-d7a29f", "xn--yalh-85da", "xn--yh-jia7671aa",
		"xn--yao-lob34x", "xn--yho-mla56g", "xn--yho-ula0g",
		"xn--yah-ina3239a", "xn--yah-czc4575a", "xn--yah-czc4475a",
		"xn--yaiho-mye", "xn--yah-e7a32e", "xn--yho-9ka08g",
		"xn--yho-e7a18d", "xn--yho-tvd5734a", "xn--yaih-tx5aa",
		"xn--aho-tna616a", "xn--ah00-kfc", "xn--yah-una413b",
		"xn--yho-hdb0328a", "xn--yh-gkaa6809a", "xn--yho-ula56n",
		"xn--yho-ela57g", "xn--yalh-8qaa", "xn--yao-czc77s",
		"xn--yho-qla1f", "xn--yho-gsb18v", "xn--aoo-vxc50e",
		"xn--yao-lob806t", "xn--yah-czc4575a", "xn--yho-ula55g",
		"xn--yaiho-uce", "xn--aho-koa719c", "xn--aho-koa619c",
		"xn--yao-see9853a", "xn--aho-bzc17d", "xn--yho-gsb612b",
		"xn--yho-ila0249a", "xn--aho-e7a06x", "xn--ah-2vc068xa",
		"xn--yoo-6cd9d", "xn--yho-ela5249a", "xn--yho-gsb08v",
		"xn--yah-7ge4843a", "xn--yah-e7a854a", "xn--hoo-8ka844b",
		"yalihoo", "yailhoo", "xn--yao-czc77s",
		"xn--yah0-8qa", "xn--yho-9oa61m", "xn--yho-9ka18g",
		"xn--aho-koa6829a", "xn--yho-hdb080b", "xn--yho-ted7884a",
		"xn--yho-gsb512b", "y-ahoo", "ya-hoo",
		"yah-oo", "yaho-o", "yazhoo",
		"yahioo", "yahnoo", "yahjoo",
		"yahoko", "yashoo", "yahoio",
		"yauhoo", "yahuoo", "yaghoo",
		"yahzoo", "yah0oo", "yahboo",
		"yahloo", "y1ahoo", "ysahoo",
		"yahopo", "yajhoo", "yaho9o",
		"yahkoo", "yzahoo", "ywahoo",
		"yanhoo", "yayhoo", "yaholo",
		"yahgoo", "yyahoo", "yahyoo",
		"yabhoo", "yahpoo", "ya2hoo",
		"yqahoo", "yawhoo", "yaho0o",
		"y2ahoo", "yah9oo", "yaqhoo",
		"yaoo", "ahoo", "yaho",
		"yhoo", "yahhoo", "yaahoo",
		"yahlo", "yah9o", "yaho0",
		"gahoo", "yaho9", "yahpo",
		"uahoo", "yazoo", "yyhoo",
		"hahoo", "yaboo", "tahoo",
		"yahoi", "aahoo", "yshoo",
		"yahio", "6ahoo", "yzhoo",
		"yahop", "yahol", "yanoo",
		"yah0o", "yauoo", "7ahoo",
		"yayoo", "yagoo", "sahoo",
		"y2hoo", "ywhoo", "y1hoo",
		"ayhoo", "yhaoo", "yaoho",
		"yuhoo", "yohoo", "yahoe",
		"yahoa", "yahuo", "yahao",
		"yahou", "yaheo", "yahoocom",
	}

	return &Brand{
		Name:       name,
		Original:   original,
		Safelist:   safelist,
		Suspicious: suspicious,
	}
}

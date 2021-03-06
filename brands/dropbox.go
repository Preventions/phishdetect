// PhishDetect
// Copyright (c) 2018-2021 Claudio Guarnieri.
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

package brands

// Dropbox brand properties.
func Dropbox() *Brand {
	name := "dropbox"
	original := []string{"dropbox"}
	safelist := []string{
		"db.tt", "dropbox.com", "dropboxapi.com", "dropboxbusiness.com",
		"dropboxdocs.com", "dropboxforums.com", "dropboxforum.com",
		"dropboxinsiders.com", "dropboxmail.com", "dropboxpartners.com",
		"dropboxstatic.com", "dropbox.zendesk.com", "getdropbox.com",
	}
	suspicious := []string{
		"dropboxa", "dropboxb",
		"dropboxc", "dropboxd", "dropboxe",
		"dropboxf", "dropboxg", "dropboxh",
		"dropboxi", "dropboxj", "dropboxk",
		"dropboxl", "dropboxm", "dropboxn",
		"dropboxo", "dropboxp", "dropboxq",
		"dropboxr", "dropboxs", "dropboxt",
		"dropboxu", "dropboxv", "dropboxw",
		"dropboxx", "dropboxy", "dropboxz",
		"eropbox", "fropbox", "lropbox",
		"tropbox", "dsopbox", "dpopbox",
		"dvopbox", "dzopbox", "dbopbox",
		"d2opbox", "drnpbox", "drmpbox",
		"drkpbox", "drgpbox", "droqbox",
		"drorbox", "drotbox", "droxbox",
		"dro0box", "dropcox", "dropfox",
		"dropjox", "droprox", "dropbnx",
		"dropbmx", "dropbkx", "dropbgx",
		"dropboy", "dropboz", "dropbop",
		"dropboh", "dropbo8", "xn--rpbx-k6b999bca",
		"xn--dpbx-v0bc77c", "xn--drbx-6vd59uca", "xn--dopbx-erc87t",
		"xn--cropbox-jhd", "xn--dpbox-erc1781c", "xn--droox-fkc57w",
		"xn--clrpbx-yxac", "xn--bropox-eof", "xn--clrpbx-dmhc",
		"xn--rpbox-4ya719b", "xn--ropbox-hyc03m", "xn--droox-fkc060a",
		"xn--dopbx-xob5240c", "xn--dr0b0x-52b", "xn--dopbx-3ta16y",
		"xn--dpbx-v0bc1349b", "xn--dropb-msf2607b", "xn--diropbox-53e",
		"droplb0x", "xn--drpbx-1ta4264c", "xn--drpx-25b2584bca",
		"xn--drplbx-cmhd", "xn--rpbx-v0bc9y", "xn--ropbx-3ta011f",
		"xn--drpbx-mye4078b", "xn--drbx-cod77oca", "xn--drplbox-m0a",
		"xn--diopbox-old", "xn--dpbox-2rc74t", "xn--drpbx-mye00u",
		"xn--dobox-krc272a", "xn--dirpbox-v2c", "dlr0pb0x",
		"xn--drpbo-1ta760d", "xn--drbx-w0bb518a", "xn--dopbx-2rc932a",
		"xn--drpb-wifc6595b", "diiropbox", "xn--rpbx-5qac980f",
		"xn--drpb0x-cxa", "xn--drobx-yva822f", "xn--drobx-9dc62z",
		"xn--brpbx-mkgc", "xn--dpbox-krc0581c", "xn--dopox-krc2r",
		"xn--drbox-4ce17r", "lbroplbox", "xn--drpox-fwc2561c",
		"xn--drpbx-jid684bda", "xn--drpbx-jsf22i", "xn--dpbox-2rc525c",
		"xn--r0pbox-hyc", "xn--dlrobox-spf", "xn--rpbox-jua753e",
		"xn--drpox-7dc317a", "xn--drpbx-kua699c", "xn--rpbx-k6b57sca",
		"xn--clrpbx-k0ec", "xn--dropbo-uyd75w", "xn--dpbox-jua05y",
		"xn--dropox-e1d90w", "xn--dopbx-mua04y", "xn--iropbox-thi",
		"xn--dpbox-erc744b", "xn--ropbx-lsf1p", "xn--drpx-mgbb498b",
		"xn--dopbx-2rc6481c", "xn--drobx-9dc518a", "xn--d0pbox-w5c",
		"xn--drpbx-sce50g", "xn--dropbx-mqc54p", "xn--drplbox-bjg",
		"xn--drpbo-1ta738d", "xn--drpox-1ta960a", "xn--drpox-fwc843c",
		"xn--dr0b0x-ycf", "xn--drpox-z0e7068b", "xn--drobx-mua600d",
		"xn--drobox-52b34y", "xn--dpbx-0ndc22f", "xn--ropbx-xfg5266b",
		"xn--brpbx-kuac", "xn--dropdox-9ke", "xn--drobx-uce51g",
		"xn--dopbx-3ta98y", "xn--dpbox-nwe8978b", "diropbox",
		"xn--drbx-wkb581aca", "dropclox", "xn--d0pbox-p6c",
		"xn--drpbx-uce3959b", "xn--dropbx-f1d28k", "xn--drpbo-kye8428b",
		"xn--dopox-nwe1p", "xn--dopbx-krc945c", "xn--dropo-gwc3911c",
		"xn--drobx-myem", "xn--drpbx-okg9246b", "blropbox",
		"xn--drpbx-jid4080cda", "xn--droplbx-fjg", "xn--drbx-1ndb51g",
		"xn--dpbx-5qac86y", "xn--bopbox-35c", "xn--dropbx-f1d89u",
		"xn--dpbox-rce52f", "xn--rpbox-0kc425a", "xn--dlropbx-f1a",
		"xn--droox-lme20d", "xn--opbox-xfg6463b", "xn--drobx-3ta325c",
		"xn--dlrpbx-k0ec", "xn--drbx-rqab04l", "xn--drplbox-b1a",
		"xn--drbox-4ce88f", "xn--drpbx-kua971e", "xn--drpbox-4wb480a",
		"xn--dopbo-nwe4k", "xn--drpibox-1ni", "xn--dopbo-nwe41g",
		"xn--clropbx-5ni", "xn--drbx-mgbb2s", "xn--drplbox-e5b",
		"xn--dropbo-uyd5647c", "xn--droplbx-1mh", "xn--dopbx-erc4781c",
		"xn--dopbx-2rc224b", "xn--drpbo-pjf08l", "xn--drbx-6qab8c",
		"xn--drbx-6qab205c", "xn--drpbx-kye9078b", "xn--drobx-lme49y",
		"xn--dopibox-rgg", "xn--diropbx-epf", "xn--iropbox-f7a",
		"xn--ropbx-9dc9y", "xn--drpx-25b850aca", "xn--dirpbox-cjg",
		"xn--ropox-0kc534a", "xn--drpdx-1tac", "xn--ropbx-9dc361c",
		"xn--drpbx-uce38r", "xn--drbox-xva021e", "xn--drpbx-lsf71i",
		"xn--dpbx-vcc25qca", "br0pbox", "xn--dropx-gkc26w",
		"xn--dropx-gwc4561c", "xn--dlrpbx-k0ec", "xn--drpbx-xob779c",
		"xn--dropox-scd89g", "xn--rpbx-k6b3484bca", "xn--brobox-kza",
		"xn--drbx-wkb1475bca", "xn--dropx-9dc55e", "xn--dpbx-0cc077bca",
		"xn--drobx-lme88n", "xn--drpbx-vob880e", "xn--drobo-lme7378b",
		"xn--drpdx-jsfc", "xn--drpbx-mye8178b", "xn--drpibx-cmhd",
		"xn--dropo-ewe5l", "xn--clrpbx-yqfc", "xn--dropx-mua65v",
		"xn--dropx-mua190d", "xn--dopbo-erc713a", "xn--drolbox-sjg",
		"xn--dropbox-qjb", "xn--dpbx-0cc64qca", "xn--drobox-e1d63p",
		"xn--drpbo-7dc7391c", "xn--dropx-gwc27r", "xn--dpbox-0ta29y",
		"xn--drpb-65dc33f", "xn--drpbx-mua199c", "xn--drpbx-9dc11z",
		"xn--ropbx-0kc83w", "xn--drobx-yva050c", "xn--drpbx-kye99i",
		"xn--dobox-2rc74x", "xn--dopbx-nwe4f", "xn--drbox-kme0129b",
		"xn--drbx-1ndbm", "xn--rpbx-k6b730aca", "xn--drpbx-mua20t",
		"xn--drpb-65dc3c", "xn--droox-dwe8h", "xn--drpbx-jid0970cda",
		"xn--robox-xva62v", "xn--d0pb0x-pof", "xn--drpibx-xqfd",
		"xn--dopbo-erc780b", "xn--dropiox-fgg", "xn--drpdx-vobc",
		"xn--dropo-ewe52g", "dropdlox", "xn--dopibox-cmd",
		"xn--ropox-fkcy", "xn--dlropbx-5ni", "xn--drbox-4ce772a",
		"xn--clrpbox-f5b", "xn--drbox-kyeo", "xn--drbox-sceo",
		"xn--dpbx-vcc0264bca", "xn--rpbx-bvec1p", "xn--drbx-csa9986bca",
		"xn--bopbox-w5c", "xn--doplbox-n51c", "xn--dropibx-5ni",
		"xn--drpbx-sce8959b", "xn--drpox-fkc06w", "xn--dropbx-fxa374b",
		"xn--dropx-mua640a", "xn--rpbox-4ya590f", "xn--drpox-1ta411d",
		"xn--drpox-sce20h", "xn--dpbox-k69ar2k", "xn--drpbx-7dc490d",
		"xn--dlrpbox-c1a", "xn--drpbx-jid80jda", "xn--drbox-xva840c",
		"xn--drbx-6vd33bc", "xn--bropbo-n77b", "xn--drpb0x-j0e",
		"xn--ropbo-pjf98j", "xn--drbox-0tb3913c", "xn--drbx-6vd74lca",
		"xn--drpbx-mye49i", "xn--d0pbox-pof", "xn--dropbo-f1d3637c",
		"xn--dirpbx-yqfc", "xn--dirpbox-n0a", "xn--ropbx-mye90s",
		"xn--diopbox-dmd", "xn--drpb-3n5atrc", "xn--ropbo-uze95r",
		"xn--drobo-yva1704c", "xn--drpdox-cxa", "xn--clrpbx-k0ec",
		"xn--drpox-fwc950a", "xn--dpbox-uob917b", "xn--droplbx-5ni",
		"xn--drobx-yva221e", "xn--drplbx-j0ed", "xn--dropb-vx1bpu",
		"xn--dpbx-vcc25qca", "xn--dropbx-f1d6097c", "dllropbox",
		"xn--rpbox-jua290f", "xn--clopbox-gld", "xn--drobox-e1d21v",
		"xn--drplbox-tx4c", "xn--dpbx-lgbc617b", "xn--dobox-xva8651c",
		"xn--rpbx-zua399bca", "drop1box", "xn--dropx-xob498b",
		"xn--drolbox-rpf", "xn--drpbx-7dc8932c", "xn--dlropbo-fi4c",
		"xn--drbx-rqab9e", "xn--drbox-7dc718a", "xn--dlropbx-fjg",
		"xn--dlopbox-o51c", "xn--dpbx-vcc41yca", "xn--brpbox-4l8b",
		"xn--brpbox-cxa", "xn--ropbx-uce36t", "xn--drobo-lme55c",
		"xn--dlrpbox-bpf", "xn--diropbo-69g", "xn--lropbox-15a",
		"xn--droplox-5rd", "xn--dobox-xva68x", "xn--rpbox-4ya9934c",
		"xn--drpbx-3ta321c", "xn--droibox-52a", "xn--bropbx-fmh",
		"xn--drobo-5ce2219b", "xn--drpb0x-xqf", "dr0plbox",
		"xn--rpbox-6dc661c", "xn--r0pbox-92a", "xn--drbx-csa34pca",
		"xn--drpox-fwc242b", "xn--rpbx-v0bc361c", "xn--ropbx-0kc725a",
		"xn--doplbox-cmd", "xn--bropbx-0xa", "xn--drbx-mgbb503b",
		"xn--dopbx-3ta86y", "xn--dropb-yob8292c", "xn--drpdx-g91bc",
		"xn--dlrpbox-f5b", "xn--drbox-xva0354c", "clroplbox",
		"xn--dropbox-ueb", "xn--dropibx-f1a", "xn--drpbx-sce492a",
		"xn--rpbx-qqac10w", "xn--brpbx-g91bc", "xn--ropbx-xob38j",
		"xn--dr0pbx-fxa", "xn--drpibx-jqcd", "xn--robox-wye26k",
		"xn--dropx-ewe5g", "xn--drbx-h6d7091bca", "xn--dpbox-krc4681c",
		"xn--brpbox-xqf", "xn--drpdox-j0e", "xn--dropb-vce25g",
		"xn--drobo-lme52k", "xn--dropx-9dc009a", "xn--drbox-7dc313a",
		"xn--diropbo-1kg", "xn--droibox-rpf", "xn--rpbox-0kc707b",
		"xn--drpx-2gc256aca", "xn--dpbox-erc462a", "xn--dopbx-uce1366b",
		"xn--opbox-4ya38w", "xn--ropbo-xfg0516b", "xn--ropbx-xfg7q",
		"xn--rpbx-5qac8p", "xn--ropdox-hyc", "xn--dpbx-063au9ica",
		"xn--drpbo-uze44i", "xn--drpbo-vx1bvt", "xn--drpbx-kua3064c",
		"clropclox", "xn--drpx-1ndb11f", "xn--ropbox-e1d350b",
		"xn--rpbx-qqac573e", "xn--drpbx-kua582f", "xn--dpbox-uob64m",
		"xn--drpbx-vob6043c", "xn--drpbo-kua640d", "xn--dropbox-5gd",
		"xn--dlrpbx-yqfc", "xn--drobx-3ta04l", "xn--drpox-kua390d",
		"xn--ropbx-xfg1166b", "xn--drpbx-mkg4346b", "xn--clropbx-fjg",
		"xn--clopbox-old", "xn--droox-5ce89e", "xn--drpdox-cmh",
		"xn--drobo-xye3328b", "xn--dlropbx-fjg", "xn--dpbox-2rc3481c",
		"xn--dropibx-epf", "xn--ropbx-4ya908c", "xn--rpbox-4ya608c",
		"xn--clropbx-xx4c", "xn--diopbox-o51c", "dlropb0x",
		"xn--bopbox-355b", "xn--dropdx-0xa", "xn--diropbx-5ni",
		"xn--dobox-xva958c", "xn--bropbx-m0e", "xn--drpibx-xxad",
		"xn--drpbx-sce50g", "xn--robox-0tb57h", "xn--drpbx-vobc180a",
		"xn--drpibox-dx4c", "xn--dpbox-nwe4878b", "xn--rpbx-0ndc890a",
		"xn--rpbox-xfg8066b", "xn--drbx-csa501cca", "xn--dopbx-2rc05t",
		"xn--iropbox-tsh", "xn--drpbo-pjf47a", "ibropbox",
		"xn--drbox-sce71g", "xn--dopbx-mua2261c", "xn--drpbx-xob509a",
		"clropdox", "xn--drpb-65dc5428b", "xn--dopbx-erc0681c",
		"xn--dpbx-qqac439c", "xn--drpx-25b4684bca", "xn--clrpbox-ex4c",
		"xn--dlrpbox-cjg", "xn--drbx-csa911dca", "xn--dlopbox-dmd",
		"xn--dpbox-erc355c", "xn--drpb-65dc5428b", "xn--dirpbx-kqcc",
		"xn--dpbox-jye5475b", "xn--drpox-7dc209a", "xn--dropbox-7ya",
		"xn--dropx-gwc8661c", "xn--dopbo-erc9931c", "xn--dpbx-55dc2475b",
		"xn--dopdox-w5c", "xn--drpbx-uce00g", "xn--dpbox-nwe47j",
		"xn--drbx-6vd7232bca", "xn--rpbx-zye1570bca", "xn--dirpbox-bpf",
		"xn--drbx-6vd9332bca", "xn--drpbx-kyec", "xn--drpx-mgbb94o",
		"xn--drpb-cvec2607b", "xn--drpbx-1ta710d", "xn--rpbox-jye21s",
		"xn--diropox-5rd", "xn--drpdx-7dcc", "xn--d0pb0x-w5c",
		"xn--drpbx-kye3278b", "dropldox", "xn--ropb0x-92a",
		"xn--rpbx-5qac08v", "xn--opbox-4ya26w", "xn--clrpbox-bpf",
		"xn--drobx-yva6454c", "xn--drpox-dwe3g", "xn--clrpbox-cjg",
		"xn--drbox-kua911c", "xn--drpbx-scec", "xn--r0pb0x-92a",
		"xn--ropox-fwc224b", "xn--dopbx-3ta439c", "xn--drpbox-jqc84p",
		"xn--dpbox-uob8240c", "xn--dropbox-qf1z", "xn--drbox-xva739c",
		"xn--dopox-dwe4385b", "xn--clropbx-q0a", "xn--drpbx-mua4854c",
		"xn--drpbx-mua80j", "dlroplbox", "xn--ropbo-0kc775a",
		"xn--dopb0x-355b", "xn--brpbox-cmh", "xn--drplbx-xxad",
		"dlropdlox", "xn--ropbx-3ta573e", "xn--ropox-z0e21r",
		"xn--drpb0x-ql8b", "xn--drpb-mgbc8292c", "xn--ropox-4ya17t",
		"xn--dpbox-2rc74t", "xn--dpbx-gdc09xca", "xn--rpbox-rce66t",
		"xn--dirpbox-ux4c", "xn--lropbox-mzb", "xn--drpbx-sce2169b",
		"xn--rpbx-lgbc298d", "xn--brobox-52b", "xn--drpb-3n5a5rc",
		"xn--drpbx-3tas", "xn--drbox-1ta1f", "xn--drpbx-uce00g",
		"xn--dlrpbox-2ni", "xn--dopbx-erc87t", "xn--ropbo-0kc743b",
		"xn--dropx-00e99h", "xn--drplbox-u2c", "xn--rpbx-55dc47k",
		"xn--dpbox-6dc25c", "xn--dpbox-uob35m", "xn--drpbo-kua618d",
		"xn--dlropbo-1kg", "xn--drobx-xob2s", "xn--droplox-5bd",
		"xn--rpbox-rce101a", "xn--drpbo-sce7319b", "xn--ropbx-xwf0637b",
		"xn--drpbx-9dc3932c", "xn--drpbx-9dc008a", "xn--drobx-lme2129b",
		"xn--drpbx-9dc11z", "xn--dobox-0tb36k", "xn--lropbox-xjg",
		"xn--brpbx-jsfc", "xn--dr0pbx-fmh", "xn--dropbx-f1d0987c",
		"xn--dpbox-2rc632a", "xn--drpx-d4d0fc", "xn--dlrpbox-ux4c",
		"xn--dropdx-m0e", "xn--dpbx-0ndc1366b", "xn--drobx-ucem",
		"xn--robox-xva4k", "xn--rpbox-0kc318c", "xn--drpdx-381bc",
		"xn--drobx-5ce3859b", "xn--dirpbx-k0ec", "xn--droplbx-epf",
		"xn--diropbx-1mh", "xn--drpbo-sce52o", "xn--dropdx-7l8b",
		"xn--drobo-yva989c", "xn--ropdox-92a", "xn--drpibox-9of",
		"xn--drpx-2gc83wca", "xn--drpbx-sce2169b", "xn--drpbx-7dc2142c",
		"xn--brpbox-qpg", "xn--drpx-25b850aca", "broplbox",
		"xn--dpbx-vcc677bca", "xn--drpox-fkc237b", "xn--dpbx-gdc4854bca",
		"xn--drpbo-sce55g", "xn--drpx-d4d25rca", "xn--drpbx-scec",
		"xn--bropox-scd", "xn--dpbox-rce4366b", "dropiibox",
		"xn--drbox-sce31b", "xn--drodox-ycf", "xn--rpbox-xwf93g",
		"xn--drpbx-9dc7042c", "xn--dopbx-nwe77j", "diroplbox",
		"biropbox", "xn--dropbox-ufb", "xn--dopbo-k69at2i",
		"xn--drbx-cod7562bca", "xn--drpb0x-qpg", "xn--drpibox-9of",
		"xn--dlropbx-y2c", "xn--brpbx-scec", "xn--rpbx-zua399bca",
		"xn--dlropbx-i5b", "xn--drbox-xva45i", "xn--dpbx-l4d4791bca",
		"xn--dirpbx-rl8bc", "xn--dr0pox-eof", "xn--drpox-dwe6978b",
		"xn--drbx-1ndb11b", "xn--bopbox-p6c", "xn--robox-4ya239b",
		"xn--dpbox-uob47m", "xn--dopbx-krc7681c", "xn--dpbx-5qac04y",
		"xn--dropo-ewe7338b", "xn--dobox-krc86x", "xn--opbox-0kc5519b",
		"xn--droox-yva19z", "xn--dropx-uce11f", "xn--droplbx-q0a",
		"xn--drpx-2gc2344bca", "xn--drpbo-1ta9414c", "xn--rpbox-jye21s",
		"xn--brpbx-7dcc", "xn--dopox-erc463a", "xn--drpox-fwc07r",
		"xn--drobx-1tb996a", "diropb0x", "xn--clrpbx-5wbc",
		"xn--dropbox-xeb", "xn--dr0pb0-gfg", "xn--robox-0tb949c",
		"xn--bropbx-7l8b", "xn--drpox-dwe68j", "xn--dobox-erc09t",
		"xn--drpbx-kua9854c", "xn--drbox-kme30c", "xn--drpibox-bjg",
		"xn--bropox-stf", "dr0pib0x", "xn--drbox-0tb685b",
		"xn--dropx-ewe88j", "xn--drobx-mua91l", "xn--drpbx-vob2933c",
		"xn--drpbx-3ta32t", "xn--drpx-2gc4444bca", "xn--droplbx-fjg",
		"xn--drpbx-mkg0246b", "xn--dopox-dwel", "xn--dpbx-vcc41yca",
		"xn--drpbox-cxa674b", "xn--drpbx-vob01h", "xn--drobx-9dc113a",
		"xn--dropbox-6gc", "xn--drpox-kua409c", "xn--dlopbox-sgg",
		"xn--rpbox-6dc109c", "xn--rpbx-zye9370bca", "dir0pb0x",
		"xn--drpx-d4d40ica", "xn--ropbx-uce36t", "xn--dr0p0x-scd",
		"xn--ropox-cwcd", "xn--clropbx-hx4c", "xn--ropbx-0kc83w",
		"xn--drobx-yva050c", "xn--drpb-6qac318d", "xn--drobx-mua711c",
		"xn--dpbx-gdc408aca", "xn--drpbx-mua8954c", "xn--ropbx-mua980f",
		"xn--drpb-mgbc616c", "xn--dlopbox-gld", "xn--dpbox-krc86t",
		"xn--brpbx-kyec", "xn--drbox-4ce5959b", "xn--dlrpbx-kqcc",
		"xn--diropbx-epf", "xn--clopbox-sgg", "xn--r0pb0x-vgh",
		"xn--dopbx-mua86y", "xn--clropbo-69g", "bropbox",
		"xn--drpx-1ndb00h", "xn--dropdx-mqc", "xn--ropox-0kcd",
		"xn--clropbo-1kg", "xn--clropbox-53e", "xn--clrpbx-dxac",
		"xn--rpbx-zua7876bca", "xn--dopbx-mua74y", "xn--bropbox-5ke",
		"xn--drobox-kza783b", "xn--drpb-1ndc4319b", "xn--drbox-4ce88f",
		"xn--drpb-mgbc648b", "xn--dr0b0x-drf", "xn--ropbx-0kc017b",
		"xn--dpbox-0ta46y", "xn--drpb-6qac340d", "xn--dpbx-vcc828aca",
		"xn--drbx-1ndbm", "xn--rpbox-0ta311f", "xn--drpbx-xob7833c",
		"xn--dopox-fkc8z", "xn--drpbx-mye49i", "xn--dropbox-vnf",
		"xn--drbox-vob229a", "xn--rpbox-xfg4q", "xn--drobx-xye8078b",
		"xn--dropbox-rfb", "xn--clrobox-tjg", "xn--opbox-1rc533c",
		"xn--drpdx-scec", "xn--dobox-wye0375b", "xn--dlrpbx-5wbc",
		"xn--dropo-00e8418b", "xn--rpbox-isf4p", "xn--dopbx-lsf9554b",
		"xn--drobx-yva2354c", "xn--ropbox-ocd", "xn--dropiox-lmg",
		"xn--drpibx-ql8bd", "xn--robox-4ya128c", "xn--drpbx-kye3278b",
		"xn--dr0p0x-stf", "xn--drpdx-kuac", "xn--dopbx-nwe7878b",
		"xn--dpbx-gdc82qca", "xn--drpox-dwe0188b", "xn--opbox-1rc006b",
		"xn--ropbx-0kc618c", "xn--drpbx-7dc508a", "xn--clropbo-fi4c",
		"xn--drobx-3ta720d", "xn--r0pb0x-vug", "dilropbox",
		"xn--clrpbox-v2c", "xn--ropbx-4ya62h", "xn--clropbx-epf",
		"xn--dr0pbx-0xa", "xn--dpbox-jua5261c", "xn--dopbo-nwe6238b",
		"xn--dpbx-0ndc22f", "xn--ropox-w0ed", "xn--dropb-nua318d",
		"xn--diopbox-gld", "xn--rpbx-0ndc890a", "xn--drbx-cod36ec",
		"xn--ropbx-0kc725a", "xn--dr0pbox-9ke", "xn--dopbx-uce22f",
		"xn--rpbox-0ta40w", "dropilbox", "xn--dr0pbx-0qf",
		"xn--dlrpbx-yxac", "xn--droox-yva748c", "xn--ropibox-f7a",
		"xn--drbx-csa343bca", "xn--dpbx-gdc09xca", "xn--drpibx-j0ed",
		"xn--dpbox-nwe1f", "xn--drpbox-e1d748a", "xn--dropbox-bza",
		"xn--drpbox-e1d97k", "xn--drplbx-xqfd", "xn--dropibx-i5b",
		"xn--drpbo-7dc526b", "xn--opbox-0kc0y", "xn--ropbx-xwf6437b",
		"xn--dirpbx-dmhc", "xn--dopbx-9dc65c", "xn--drpb-1ndc22o",
		"xn--drpibx-xqfd", "xn--iropbox-xcd", "xn--drobx-lme8919b",
		"xn--drob0x-y0e", "xn--dlrpbx-yqfc", "xn--dpbox-6dc527a",
		"xn--drpb-1ndc25g", "xn--dopbx-nwe1088b", "xn--drbox-wye28i",
		"xn--bropbx-7wb", "xn--dirpbx-5l8bc", "xn--drpbo-7dc558a",
		"xn--drpb0x-xqf", "xn--drpx-1ndb11f", "xn--drpbx-vob289c",
		"xn--dpbox-0ta739c", "xn--drpbx-1ta002e", "xn--brpbox-j0e",
		"xn--dopbx-erc762a", "xn--drobx-5ce7959b", "xn--dropx-uce00h",
		"xn--drbx-h6d36qca", "xn--droox-fwc470a", "xn--ropox-dwe66l",
		"xn--drpbx-jid79rda", "xn--drbx-65dbm", "xn--dropdx-tl8b",
		"xn--drpox-z0e30t", "xn--r0pb0x-hyc", "xn--dirpbox-c1a",
		"xn--drpbx-jid083ada", "xn--drpbo-uze05t", "xn--dropx-gkc155a",
		"xn--drpx-25b69sca", "xn--dpbx-l4d6891bca", "xn--drplbox-bjg",
		"xn--doplbox-rgg", "dr0plb0x", "xn--drpbo-kye63f",
		"xn--drpx-6qab640a", "xn--dropx-9dc5w", "xn--dobox-kme52b",
		"xn--rpbx-zua559dca", "xn--drpibox-xmh", "xn--drobx-yva65i",
		"xn--ropbox-9nf", "xn--dlropox-5bd", "xn--drplbx-ql8bd",
		"xn--dobox-4ce9166b", "xn--drpbx-uce3959b", "xn--robox-0kc05w",
		"xn--drpox-fkc945a", "xn--drpox-1ta97v", "xn--dropb-bec4391c",
		"xn--drpx-77d4381bca", "xn--dropb-pkg6595b", "xn--drobox-e1d69k",
		"xn--rpbox-isf92g", "xn--robox-4ya43j", "clr0pbox",
		"xn--drpx-65db1h", "lbropbox", "xn--dpbx-qqac3461c",
		"xn--drpbx-7dc61z", "xn--ropbx-mua453e", "xn--drpx-mgbb95j",
		"xn--dlropbx-epf", "xn--rpbox-xwf3437b", "xn--dopbox-35c52j",
		"xn--dopbox-e1d2054c", "xn--dirpbox-cjg", "xn--drobx-yva05s",
		"xn--clrobox-d2f", "xn--robox-0kc935a", "xn--drplbox-xmh",
		"xn--drplbx-4wbd", "xn--opbox-mwe09s", "xn--dpbox-2rc914b",
		"xn--drobx-1tb99e", "xn--diropbx-f1a", "xn--dopbx-krc3581c",
		"xn--lropbox-tsh", "xn--dlopbox-old", "xn--drpibx-qpgd",
		"xn--dpbox-erc7581c", "xn--brpbx-scec", "xn--droibox-c2f",
		"xn--drbox-kme6919b", "xn--dopbx-erc054b", "xn--drbox-kme30c",
		"xn--drpx-2gc83wca", "xn--drpb-1ndc4319b", "xn--drop0x-sxc",
		"xn--drop0x-stf", "xn--drpibox-e5b", "xn--bropbx-tl8b",
		"xn--droox-lme31b", "xn--rpbox-jua1q", "xn--rpbox-0kc1312c",
		"xn--clropbx-1mh", "xn--drpibox-tx4c", "xn--drpox-7dc7w",
		"xn--drpox-kye3h", "xn--rpbx-zef1210bca", "xn--drpbx-9dc389b",
		"xn--drpbx-uce7069b", "xn--dlropbx-hx4c", "xn--dopbx-mua319c",
		"xn--bropbo-gsf", "xn--drpbx-sce50g", "xn--roplbox-xcd",
		"xn--dpbx-0cc80yca", "xn--dropx-xob507b", "xn--dopbx-2rc2381c",
		"xn--dropibo-fi4c", "xn--drodox-drf", "xn--drobx-1tb996a",
		"xn--dropb-nua5214c", "xn--dopbx-2rc05t", "xn--drbx-wkb3575bca",
		"xn--clropox-lmg", "xn--bopbox-pof", "xn--dlrpbx-rl8bc",
		"xn--drbox-xva84s", "xn--rpbx-55dc90s", "xn--rpbx-5qac453e",
		"xn--dopox-erc9r", "xn--dropb-vx1byt", "xn--dobox-0tb24k",
		"bropb0x", "xn--bropbo-gfg", "xn--opbox-jrc126b",
		"xn--drpbx-xob50h", "xn--dopbx-krc17t", "xn--drpox-fwc950a",
		"xn--opbox-drc363c", "xn--dobox-erc57x", "xn--dropx-3ta760a",
		"xn--drbx-csa760eca", "xn--dropb0-gfg", "xn--drpbx-kuac",
		"xn--rpbx-55dc90s", "xn--dopbox-p6c10j", "xn--drbox-4ce1859b",
		"xn--dopbx-okg3543b", "xn--drobx-xob029a", "xn--rpbox-0kc53w",
		"xn--dpbx-l4d49hca", "xn--dopibox-fld", "xn--ropox-fkc219b",
		"xn--drpbx-vobc", "xn--dr0box-kza", "xn--brpbx-kyec",
		"xn--dr0pbx-m0e", "xn--dpbx-l4d0ec", "xn--drpb-w0bc226b",
		"xn--drpbx-xob497b", "xn--dropx-mua209c", "xn--drpbx-mua201c",
		"xn--dr0box-y0e", "xn--drpbx-vob019a", "xn--brobox-y0e",
		"xn--dropo-gkc106a", "xn--dropox-zgh", "xn--dlrpbox-ex4c",
		"dropibox", "xn--dropbox-oza", "xn--r0pbox-vug",
		"xn--dlropbx-1mh", "xn--drpbx-kyec", "xn--drpbx-1ta710d",
		"xn--dropibx-xx4c", "xn--drpbx-kyec", "xn--dropdx-0qf",
		"xn--drpdx-kyec", "xn--rpbox-uob598d", "xn--drbx-csa343bca",
		"xn--dopbo-krc2931c", "xn--dropbx-7wb180a", "xn--ropbo-4ya1493c",
		"xn--dropox-lyc", "xn--ropox-4yad", "xn--ropbx-4ya299d",
		"xn--drbox-xva840c", "xn--drbox-sce31b", "xn--clrpbx-rpgc",
		"xn--dropb-vce25g", "xn--diropox-fgg", "xn--clopbox-dmd",
		"xn--diropbx-xx4c", "xn--drbx-6qab600d", "xn--dropb-nye5428b",
		"xn--dpbx-vifc3543b", "xn--dpbx-l4d0ec", "xn--rpbx-lgbc38j",
		"xn--ropox-4ya16y", "xn--drpx-77d09pca", "xn--drpb-w0bc258a",
		"xn--dopdox-355b", "xn--drpbox-e1d97k", "xn--drpx-rqab760a",
		"xn--rpbox-jua38v", "xn--dropo-gkc3861c", "dropib0x",
		"xn--drpibox-u2c", "xn--drpbx-7dc889b", "xn--dropx-xob95j",
		"xn--dpbox-erc57t", "xn--drpbx-kye9078b", "xn--dlrpbox-cjg",
		"xn--drpb-w0bc4391c", "xn--clrobox-bdc", "xn--dopb0x-pof",
		"xn--rpbox-uob061d", "xn--drbx-65dbm", "xn--drpbo-vob948b",
		"xn--drbx-h6d51hca", "xn--drpibox-m0a", "xn--robox-kme66p",
		"xn--dropx-00e3268b", "xn--drbox-1ta920d", "xn--dropx-gkc8512c",
		"xn--drbox-0tb685b", "xn--drpbx-uce00g", "xn--drbx-cod52yca",
		"xn--dpbx-063a48ica", "xn--drbx-6qab91l", "xn--lropbox-f7a",
		"xn--drobo-5ce01o", "xn--ropox-dwe10t", "xn--dpbox-lkg6543b",
		"xn--droox-fkc465a", "xn--dropdo-gsf", "xn--dpbx-qqac16y",
		"xn--dropdo-gfg", "xn--dpbx-l4d24rca", "xn--dropbox-ynf",
		"xn--dlropox-5rd", "xn--drobo-5ce04g", "xn--ropibox-xcd",
		"xn--drpx-w0bb009a", "xn--brpbox-xqf", "xn--droox-1tb586b",
		"xn--dopox-krc861a", "xn--drpb0x-j0e", "xn--drpbx-uce982a",
		"xn--drpbx-kuac944b", "xn--drpbx-jsf6257b", "xn--drpbo-vob916c",
		"xn--rpbox-4ya32h", "xn--robox-0kc530a", "xn--drobx-ucem",
		"dropd0x", "droplibox", "xn--dlrpbx-rpgc",
		"xn--dobox-erc972a", "xn--drpox-vob707b", "xn--dopox-fkc0819b",
		"xn--ropox-0kc97b", "xn--ropbox-oxc", "xn--drpdox-jqc",
		"xn--drolbox-9cc", "xn--dr0pb0x-9ke", "xn--dopbo-krc013a",
		"xn--ropbx-3ta10w", "xn--drbx-csa1196bca", "xn--drbx-wkb581aca",
		"xn--bropbox-9ke", "dliropbox", "xn--drbx-w0bb62z",
		"xn--ropbox-otf", "xn--dpbox-krc645c", "xn--drbox-kme29y",
		"xn--drpibox-bjg", "xn--dopbx-krc17t", "xn--drob0x-ycf",
		"xn--drpbx-3ta210d", "xn--dropx-9dc117a", "xn--rpbox-4ya989d",
		"xn--drpbx-scec", "xn--drpbx-3ta9164c", "xn--drpibx-cxad",
		"xn--dropx-mye1h", "xn--drbx-1ndb51g", "xn--drpbx-sce8959b",
		"xn--dropx-ewe8978b", "xn--robox-4ce15t", "xn--drpox-z0e1268b",
		"xn--dr0pbx-m0e", "xn--brpbox-4wb", "xn--drpdx-scec",
		"xn--rpbox-jye77k", "xn--brpbox-ql8b", "xn--drpbx-okg5146b",
		"xn--rpbox-4ya3144c", "xn--clrpbox-c1a", "xn--clrobox-62a",
		"xn--drpox-sce31f", "xn--dpbx-5qac74y", "xn--drobo-1tb0272c",
		"xn--rpbx-bvec62g", "xn--drpbx-381bqa", "xn--drpbx-kye50u",
		"xn--bropbx-tpg", "xn--dirobox-bdc", "xn--drbox-xva4454c",
		"xn--dr0pbx-7l8b", "xn--dopbo-krc080b", "xn--clropox-fgg",
		"xn--dropx-uce11f", "xn--drpibx-4l8bd", "xn--drplbx-qpgd",
		"xn--drbox-0tb578d", "xn--drbox-1ta24l", "xn--dobox-0tb53k",
		"xn--dpbx-55dc2475b", "xn--ropibox-tsh", "xn--dpbox-jua619c",
		"xn--drobx-uce11b", "xn--drbx-csa19gca", "xn--drpdox-xqf",
		"xn--drob0x-kza", "bropdox", "xn--rpbox-xfg2266b",
		"xn--dr0pbx-7wb", "xn--rpbox-0kc425a", "xn--dopbx-xob05m",
		"xn--brpbx-381bc", "xn--dpbx-lgbc17m", "xn--drpx-d4d0fc",
		"xn--dropbo-uyd735a", "xn--drobo-xye1b", "xn--dlropox-lmg",
		"xn--dropb-uze74i", "diropibox", "xn--drobo-xye12f",
		"xn--drpbo-sce7319b", "xn--dropb-pjf77a", "xn--droibox-sjg",
		"xn--drpb-yne3021bca", "xn--dopbx-mye2475b", "xn--ropox-0kc426a",
		"xn--rpbx-v0bc898c", "xn--dlropbx-xx4c", "xn--dropb-nua340d",
		"xn--ropbox-e1d398b", "xn--rpbx-zua13oca", "xn--clopbox-o51c",
		"xn--cropbox-njb", "xn--drbx-6vd33bc", "xn--drpbx-1tac",
		"dlropbox", "xn--clrpbox-bpf", "xn--drbox-vob4s",
		"xn--drpox-fkc2412c", "dlropdox", "xn--dropbox-2gd",
		"xn--drpbx-kua701c", "xn--drbx-1ndb11b", "xn--drbox-wye88t",
		"xn--dropibo-1kg", "xn--drpb0x-4wb", "xn--ropox-4ya609c",
		"xn--robox-kme10x", "xn--dpbx-qqac98y", "xn--drpbox-e1d111c",
		"xn--dopbo-pjf2294b", "xn--drpox-kua840a", "xn--dropo-uze1c",
		"xn--dlropbox-53e", "xn--dopbx-xob617b", "xn--dpbox-krc752a",
		"xn--dr0b0x-y0e", "xn--drplbx-j0ed", "xn--rpbx-55dc47k",
		"xn--drbox-0tb79e", "xn--dropb-uze35t", "xn--dropb-nye33f",
		"xn--dropx-gkc048c", "xn--drpb-65dc3c", "xn--dopibox-n51c",
		"xn--rpbox-xwf7537b", "xn--drbx-wkb749aca", "xn--ropbx-4ya908c",
		"xn--dopbx-k69au2k", "xn--dopbx-krc062a", "xn--drpox-7dc75e",
		"xn--drpb-86d3681bca", "xn--diropbx-i5b", "xn--clropbx-epf",
		"xn--dpbox-0ta17y", "xn--ropbx-uce890a", "xn--drpb-65dc33f",
		"xn--dropdx-tpg", "xn--dpbx-0cc4164bca", "xn--drobx-xob503b",
		"xn--brpbox-jqc", "clr0pb0x", "dr0pibox",
		"xn--dropb-nye3c", "xn--dropb-pjf38l", "xn--ropbx-xob751d",
		"xn--drpbx-mua082f", "xn--dropibx-hx4c", "xn--dopbx-9dc77c",
		"xn--dropb-4ta438d", "xn--dropo-00e63e", "xn--dropdx-m0e",
		"xn--drpbx-7dcc", "xn--drbx-rqab720d", "xn--drobx-1tb5913c",
		"xn--dopbo-uze2965b", "xn--clropox-5bd", "xn--dlrobox-bdc",
		"xn--drpb0x-cmh", "xn--dropibo-69g", "xn--bropbox-5kee",
		"xn--drobx-5ce09f", "xn--drbox-kua0d", "xn--drbox-vob703b",
		"xn--drpbo-uze8668b", "xn--dropbox-jhd", "xn--dobox-4ce01f",
		"xn--drobx-5ce37r", "xn--droox-yva630d", "xn--drplbox-9of",
		"xn--clrpbox-ymh", "xn--dropbox-3gc", "xn--droplbx-f1a",
		"xn--rpbx-zef9010bca", "xn--d0pb0x-p6c", "xn--dopox-z0e5465b",
		"xn--rpbox-0kc7112c", "dropdox", "xn--diropbx-hx4c",
		"xn--drpbx-kye99i", "xn--drpbo-kye8428b", "xn--lropbox-4pf",
		"xn--drolbox-52a", "xn--clrpbx-5l8bc", "xn--drpbx-7dc508a",
		"xn--ropbx-4ya890f", "xn--drpbo-kye63f", "xn--rpbox-0ta873e",
		"xn--ropbo-4ya958c", "xn--dropox-sxc85m", "xn--rpbx-zua9976bca",
		"xn--drpbo-kye6c", "xn--dr0box-drf", "xn--ropbo-4ya926d",
		"xn--diropbo-fi4c", "xn--dlrpbox-v2c", "xn--dropx-gkc155a",
		"xn--dopbx-xob34m", "xn--dopbx-krc344b", "xn--clrpbx-kqcc",
		"xn--dpbox-rce4366b", "xn--ropox-4ya717c", "xn--drobx-yva939c",
		"xn--dpbx-5qac319c", "xn--drobx-yva939c", "xn--dopbx-uce22f",
		"xn--bropbx-mqc", "xn--dropx-00e9068b", "xn--dropx-gkc4412c",
		"xn--dopbx-k69a82k", "xn--clropox-5rd", "xn--drbox-xva739c",
		"xn--dpbox-6dc08c", "xn--dr0pbx-0qf", "xn--dropbox-rza",
		"xn--bropox-sxc", "xn--rpbx-zua132bca", "xn--drbox-wye6078b",
		"xn--rpbx-zua700dca", "xn--drpbo-pjf8987b", "xn--drplbx-4l8bd",
		"xn--drbox-0tb796a", "xn--droplbo-fi4c", "xn--dopox-fwc0968b",
		"xn--drpbx-3ta5064c", "xn--droplbx-i5b", "xn--drpx-w0bb117a",
		"xn--drobo-yva957d", "xn--dpbox-erc462a", "d1ropbox",
		"xn--dopbx-2rc825c", "xn--drpx-6qab65v", "xn--drolbox-c2f",
		"xn--rpbx-0ndc36t", "xn--dopox-fkc62a", "xn--droox-1tb04m",
		"xn--droplbx-y2c", "xn--drobx-1tb885b", "xn--drpx-65db1h",
		"xn--bropbx-0qf", "xn--drpbx-jsf0457b", "xn--ropbx-mye47k",
		"xn--drbx-wkb909cca", "xn--drpbx-lsf1257b", "xn--dpbox-rce52f",
		"xn--ropox-xfgd", "xn--dopb0x-p6c", "xn--dr0pbx-mqc",
		"xn--rpbx-zua132bca", "xn--dirpbox-ex4c", "xn--dpbx-gdc82qca",
		"xn--diropbx-fjg", "xn--drbx-wkb749aca", "xn--drpbx-1ta821c",
		"cliropbox", "xn--dropibx-fjg", "xn--dlropbx-q0a",
		"xn--dpbox-6dc4349b", "xn--drobx-mua205c", "xn--dlrpbx-dxac",
		"xn--dopdox-p6c", "xn--dropb-nye5428b", "dr0pdox",
		"xn--rpbox-6dc2z", "xn--opbox-drc826b", "xn--rpbox-0kc53w",
		"xn--ropox-z0e77j", "xn--drpbo-kua8214c", "xn--dropox-zug",
		"xn--brobox-drf", "xn--opbox-4ya7341c", "xn--drpox-fwc6661c",
		"xn--dopbx-xob17m", "xn--dpbox-2rc9281c", "xn--dropx-ewe2188b",
		"xn--dopb0x-35c", "xn--drpdox-j0e", "xn--drbox-0tb796a",
		"xn--drobx-lme50c", "xn--dpbx-gdc257bca", "xn--r0pbox-vgh",
		"xn--dropbox-nf1z", "xn--dr0pbx-tpg", "xn--ropox-fkc746c",
		"xn--diopbox-sgg", "xn--dlropbo-69g", "xn--drpb-rqac438d",
		"xn--droplox-fgg", "xn--dpbx-lgbc34m", "xn--dpbox-krc86t",
		"xn--dpbox-nwe1f", "xn--drpbx-7dc61z", "xn--drplbox-1ni",
		"xn--dopox-2rc633a", "xn--drpbx-381bc", "xn--drpbx-3ta321c",
		"xn--drbox-sce71g", "xn--drpx-77d24gca", "xn--dropb0-n77b",
		"xn--drpbx-vob997b", "xn--dropbox-mhd", "dr0pd0x",
		"xn--dropx-gwc160a", "xn--drpbx-kye50u", "xn--drpbx-xob380e",
		"xn--dropb-4ta6414c", "xn--dropx-3ta211d", "ciropbox",
		"xn--dropibx-1mh", "xn--rpbx-qqac9r", "xn--dpbox-6dc95c",
		"dlropibox", "xn--brpbox-j0e", "xn--dobox-2rc26t",
		"xn--dobox-kme4326b", "xn--dr0pox-sxc", "xn--rpbox-rce101a",
		"xn--dpbox-nwe08u", "xn--drpb0x-jqc", "xn--dobox-xva41y",
		"xn--dpbox-jua34y", "xn--clrpbx-rl8bc", "xn--dpbx-qqac86y",
		"xn--ropbx-3ta9r", "xn--drpx-2gc67oca", "diropdox",
		"xn--drpbx-uce7069b", "xn--ropbx-9dc898c", "xn--drpb-rqac460d",
		"xn--drpibx-4wbd", "xn--dlrobox-62a", "xn--brpbox-xxa",
		"xn--dpbox-2rc632a", "xn--drpbx-kyec", "xn--dropdx-0qf",
		"xn--drpx-77d2281bca", "xn--dopbx-2rc932a", "xn--droox-fwc07v",
		"xn--dropx-00e50t", "xn--drbx-wkb169bca", "xn--drbox-kua405c",
		"xn--drbx-mgbb908b", "xn--dlropox-fgg", "xn--droplox-lmg",
		"xn--drpbx-mua471e", "xn--drpbx-xob509a", "xn--dlrobox-tjg",
		"xn--drplbx-xqfd", "xn--ropb0x-vug", "xn--dobox-2rc152a",
		"xn--dr0pox-scd", "xn--rpbox-rce66t", "xn--droox-yva10v",
		"xn--dropb-vce4319b", "xn--dropdx-7wb", "br0pb0x",
		"xn--drpbx-mua199c", "xn--drbox-0tb9713c", "xn--dpbox-krc044b",
		"xn--droox-fwc58r", "xn--dropibox-53e", "xn--dpbx-0cc2064bca",
		"dropidox", "xn--dopbox-e1d60u", "xn--drpbx-lsf5357b",
		"xn--dobox-nwe6g", "xn--drpbx-kua699c", "xn--ropbox-92a832b",
		"xn--dropox-e1d39t", "xn--dpbx-v0bc227a", "xn--drpbx-1ta43j",
		"dr0pb0x", "xn--drpbx-vob019a", "xn--lropbox-thi",
		"xn--drpx-d4d6991bca", "xn--clropbx-i5b", "xn--drpbx-581bl",
		"xn--clrpbox-cjg", "xn--drpbx-mye4078b", "xn--drpox-fkc06w",
		"xn--dirpbx-rpgc", "xn--dirpbx-yqfc", "dropllbox",
		"xn--drpbx-mye00u", "xn--drpbo-jsf5607b", "xn--drbox-wye2968b",
		"xn--drpox-kye3h", "xn--droox-1tb694b", "xn--drodox-52b",
		"xn--brobox-ycf", "xn--rpbx-zef7nc", "dir0pbox",
		"c1ropbox", "xn--dopox-erc571a", "xn--drobx-uce51g",
		"xn--robox-wye79r", "xn--drpbx-jid80jda", "xn--drpbx-kua70t",
		"xn--bropbx-0qf", "xn--dlrpbox-bpf", "xn--ropbx-lsf62g",
		"ibropibox", "xn--clrpbox-n0a", "xn--dropbx-f1d28k",
		"xn--rpbox-4ya719b", "xn--dropb-4ta460d", "xn--ropbx-mye47k",
		"xn--ropbx-mua08v", "xn--droox-5ce78g", "xn--dlrpbox-ymh",
		"xn--drobx-1tb1813c", "xn--drpbo-kye6c", "xn--ropbx-4ya6144c",
		"xn--drbx-h6d5981bca", "xn--drpb-86d5781bca", "xn--drbox-xva622f",
		"xn--dopox-2rc1p", "xn--drpox-fkc945a", "xn--drpb-yne32ac",
		"xn--ropb0x-vgh", "dlr0pbox", "xn--drpbo-pjf4887b",
		"xn--drpx-25b69sca", "xn--drpb-86d38gca", "xn--drobx-xye09t",
		"xn--drpbo-uze4568b", "xn--roplbox-tsh", "xn--dropbx-f1d89u",
		"xn--drpb-yne5121bca", "xn--drbox-sceo", "xn--drpdox-4l8b",
		"xn--drpbx-jid79rda", "xn--dobox-0tb7120c", "xn--dirpbox-f5b",
		"xn--dpbx-lgbc05m", "xn--dobox-xva39x", "xn--bropbx-fxa",
		"xn--opbox-0kc3x", "clropb0x", "xn--drplbox-9of",
		"xn--dpbx-bvec9554b", "xn--drbox-vob118b", "xn--dr0pb0-n77b",
		"xn--drop0x-scd", "xn--dpbx-0cc80yca", "xn--drpbox-e1d7887c",
		"xn--dirpbox-2ni", "xn--ropibox-thi", "xn--drop0x-eof",
		"xn--drpbx-sce50g", "xn--rpbox-uob68j", "xn--dr0p0x-eof",
		"xn--ropbx-4ya02r", "xn--drpbo-sce52o", "xn--dropx-mye1h",
		"xn--dropox-e3a", "xn--opbox-0kc644a", "xn--drpox-vob15o",
		"xn--drob0x-52b", "xn--drpb0x-xxa", "xn--drpdox-xqf",
		"xn--drpdox-xxa", "xn--dropbox-j95a", "xn--drpbx-uce38r",
		"xn--drobx-3ta9e", "xn--dobox-krc38t", "xn--drbx-mgbb029a",
		"xn--dpbox-isf2654b", "xn--drpbx-uce982a", "xn--drpx-w0bb5w",
		"xn--drpx-rqab77v", "xn--robox-4ya713c", "xn--drobx-xob908b",
		"xn--dropb-pjf1097b", "xn--drobx-myem", "xn--diropbx-q0a",
		"xn--droibox-9cc", "xn--drpx-2gc006bca", "xn--dropb-uze1768b",
		"xn--drpbx-mye8178b", "xn--drpox-dwe29u", "xn--dpbox-0ta6461c",
		"xn--ropox-ckcd", "xn--drbx-csa501cca", "xn--rpbox-4ya608c",
		"xn--drpox-vob16j", "xn--ropb0x-hyc", "xn--drodox-kza",
		"xn--dropx-gwc27r", "dropdiox", "xn--dr0pox-stf",
		"xn--dopbx-nwe38u", "xn--drpx-d4d4891bca", "xn--brpbx-vobc",
		"xn--droplbx-xx4c", "xn--dopbx-nwe4f", "xn--drpbox-e1d59u",
		"xn--clropbx-y2c", "xn--dirobox-d2f", "xn--dopbx-erc762a",
		"xn--drpox-z0e79h", "xn--drpbo-sce55g", "xn--drbx-rqab831c",
		"xn--drpox-fkc6512c", "xn--dopbx-9dc227a", "xn--drpbx-1ta603f",
		"xn--drbx-cod36ec", "xn--droplbox-53e", "xn--dopbx-9dc1349b",
		"xn--clrpbox-2ni", "xn--ropbx-0kc0212c", "xn--drpx-2gc67oca",
		"xn--dr0pbo-gsf", "diropdiox", "xn--dopbo-2rc1731c",
		"xn--dopbo-2rc982a", "xn--dropx-ewe49u", "xn--drpox-sce31f",
		"xn--droplbo-69g", "xn--ropdox-vgh", "xn--dopbx-3ta3461c",
		"xn--dopbx-erc655c", "clropibox", "xn--drpb-86d13qca",
		"xn--dirpbx-yxac", "xn--dr0p0x-sxc", "xn--drpbx-scec",
		"xn--d0pb0x-35c", "xn--dropibx-y2c", "xn--dropb0x-9ke",
		"xn--drpbx-kua701c", "xn--d0pb0x-355b", "xn--dopox-fwc18z",
		"xn--drpox-fkc838c", "xn--dlrpbx-dmhc", "cllropbox",
		"xn--drpb-yne17jca", "xn--drobx-xye48i", "xn--clrpbx-yqfc",
		"xn--drbox-0tb967c", "xn--dopbo-2rc950b", "xn--dropdo-n77b",
		"xn--dopibox-nld", "xn--droplbo-1kg", "xn--dropx-3ta329c",
		"xn--drpbx-1ta82t", "xn--robox-xva003e", "xn--drpx-6qab190d",
		"xn--dr0pbo-gfg", "xn--ropbx-mye90s", "xn--drpbx-xob1043c",
		"xn--d0pbox-355b", "xn--dropdx-fxa", "xn--dropx-xob94o",
		"xn--drpbox-e1d3097c", "xn--drpbx-1ta0164c", "xn--drpibox-b1a",
		"xn--drpbx-3ta103f", "xn--clrobox-spf", "xn--drpibx-j0ed",
		"xn--dropx-3ta77v", "xn--ropbx-xob298d", "xn--ropbo-xwf5886b",
		"xn--rpbx-lgbc751d", "xn--dlrobox-d2f", "xn--drbx-w0bb113a",
		"xn--dpbx-0cc228aca", "xn--ropbx-4ya029b", "xn--dpbox-jua17y",
		"xn--dropbox-9kea", "xn--robox-0tb486d", "xn--drpbx-3ta210d",
		"xn--drpx-25b020cca", "xn--drpbo-vx1b9t", "xn--opbox-jrc653c",
		"xn--drpox-kua85v", "xn--drpx-mgbb507b", "xn--dopbx-mye2475b",
		"xn--dropb-bec226b", "xn--drpx-rqab329c", "xn--dropbx-f1d411c",
		"xn--opbox-4ya827c", "xn--rpbx-zye70fca", "xn--rpbox-4ya71r",
		"xn--drpox-1ta529c", "xn--dropx-gwc053c", "xn--clropbx-f1a",
		"xn--drpbx-3ta92j", "xn--dropbo-f1d534a", "xn--roplbox-f7a",
		"xn--drobx-1tb177c", "xn--drobx-5ce09f", "xn--ropbx-xwf24g",
		"xn--dirpbx-k0ec", "xn--dropb-bec258a", "xn--ropox-xwfd",
		"xn--drpbx-9dc980d", "xn--dpbx-v0bc65c", "xn--rpbx-qqac011f",
		"xn--droplbx-epf", "xn--dirpbx-5wbc", "xn--drob0x-drf",
		"xn--drpox-vob698b", "xn--dropbo-f1d55v", "xn--dropiox-5bd",
		"xn--drpx-6qab209c", "xn--dropb-nye3c", "xn--dropx-ewe5g",
		"xn--dopbx-krc062a", "xn--ropbo-uze42k", "xn--dropx-gkc26w",
		"xn--dropbx-f1d058a", "xn--dpbx-lgbc5240c", "xn--drobo-1tb836b",
		"xn--ropbx-uce890a", "xn--clrpbox-ux4c", "xn--ropbx-4ya029b",
		"xn--opbox-xwf1834b", "xn--drpbx-g91bc", "xn--drpbo-vob1392c",
		"xn--dr0b0x-kza", "xn--dirobox-spf", "xn--drbox-7dc82z",
		"xn--dirpbox-ymh", "xn--rpbx-0ndc36t", "xn--opbox-4ya55w",
		"xn--dr0box-ycf", "xn--dpbox-jye5475b", "xn--diropbx-y2c",
		"xn--drpbox-e1d59u", "xn--dopox-fkc174a", "xn--dopbox-w5c33j",
		"xn--drodox-y0e", "xn--dropb-vce22o", "xn--drpx-rqab211d",
		"xn--dpbx-vcc8064bca", "xn--rpbx-k6b150bca", "xn--dpbx-gdc6954bca",
		"xn--drobx-3ta831c", "xn--dr0pbo-n77b", "xn--drpb-6qac5214c",
		"xn--rpbx-k6b57sca", "xn--drbox-1ta041c", "xn--dopb0x-w5c",
		"xn--drpdox-ql8b", "xn--robox-xva530f", "xn--dropo-gkc173b",
		"xn--clropbx-fjg", "xn--dirpbox-bpf", "xn--dpbx-0ndc1366b",
		"xn--ropbo-0kc9561c", "xn--drpbx-jsfc", "xn--dopbx-9dc94c",
		"xn--dropx-gkc437b", "xn--drobo-1tb804c", "xn--rpbx-k6b1384bca",
		"xn--dopdox-pof", "xn--ropbx-4ya2044c", "xn--ropbo-pjf45c",
		"xn--drbx-rqab325c", "xn--drpb0x-4l8b", "xn--dropb-nye33f",
		"xn--dpbox-erc57t", "xn--drpdox-4wb", "xn--dirpbx-dxac",
		"xn--drobx-mua8c", "xn--dropb-uze7568b", "xn--bropbx-m0e",
		"xn--dlrpbox-n0a", "xn--drbox-1ta525c", "xn--drpb-rqac6414c",
		"xn--diropbx-fjg", "xn--ropbx-mua8p", "xn--drpbx-1ta821c",
		"xn--drbx-cod9662bca", "xn--drbx-6qab711c", "xn--drpbx-sce88r",
		"xn--dpbox-k69a52k", "xn--drpbo-mkg9595b", "xn--dr0pbx-tl8b",
		"xn--rpbox-0ta2s", "xn--dropb-vce4319b", "xn--dropb-yob648b",
		"xn--drbox-kyeo", "xn--drbox-kua800d", "xn--dlrpbx-5l8bc",
		"xn--drpox-fwc07r", "xn--diropox-lmg", "xn--drpdx-mkgc",
		"xn--dirobox-62a", "xn--dopbx-uce1366b", "xn--drpbx-sce88r",
		"xn--dirobox-tjg", "xn--drplbx-jqcd", "xn--dropo-gwc111a",
		"xn--dpbx-v0bc94c", "xn--drpbx-1tax", "xn--drbx-wkb58dc",
		"xn--drobx-5ce972a", "xn--drpox-sce20h", "xn--dropo-gwc188a",
		"xn--lropbox-44a", "xn--dopdox-35c", "xn--drpbx-1tac374b",
		"xn--drplbox-dx4c", "xn--dropbx-0xa944b", "xn--dopox-krc753a",
		"xn--dropx-gwc442b", "xn--dropb-yob616c", "xn--drpbx-xob497b",
		"xn--brpbx-1tac", "xn--drpx-25b270bca", "xn--drplbx-cxad",
		"xn--drobx-1tb778d", "xn--rpbx-k6b730aca", "xn--ropox-fwc751c",
		"xn--dropb-vce22o", "xn--drobx-1tb885b", "xn--drbox-kme68n",
		"xn--dr0pb0-gsf", "xn--dropiox-5rd", "xn--d0pbox-35c",
		"xn--drobx-lme50c", "xn--dropb-pjf7887b", "xn--dropx-uce00h",
		"xn--dopox-2rc741a", "xn--drpbx-sce492a", "xn--drpbx-vob997b",
		"xn--dopox-fkc50a", "xn--droox-xye8f", "xn--dpbx-0cc64qca",
		"xn--dropx-gwc160a", "xn--dlropbx-epf", "xn--roplbox-thi",
		"xn--diropox-5bd", "xn--doplbox-nld", "xn--dropibx-fjg",
		"xn--drpbx-kua31j", "xn--drpbx-uce00g", "xn--dropibx-epf",
		"xn--drpdox-qpg", "xn--drpb-1ndc25g", "xn--drobx-uce11b",
		"xn--dpbox-krc752a", "xn--droplbx-hx4c", "xn--drpbx-7dcc54p",
		"xn--drpbox-xxa254b", "xn--drpbx-mkgc", "xn--doplbox-fld",
		"xn--rpbox-jye77k", "xn--dropdx-fmh", "xn--dpbx-5qac2261c",
		"xn--ropdox-vug", "xn--dropibx-q0a", "xn--drpbx-3ta591e",
		"xn--robox-4ce680a", "xn--drpbx-9dc008a", "xn--ropbx-0kc4312c",
		"xn--lropbox-xcd", "bropibox", "xn--drpb-1ndc22o",
		"xn--opbox-0kc10a", "xn--dropb0-gsf", "xn--rpbx-zua97fca",
		"xn--drpx-w0bb55e", "xn--lropbox-txa", "xn--dropbox-m95a",
		"xn--dr0box-52b", "xn--droox-1tb05h", "xn--drpox-dwe3g",
		"xn--dropbox-njb", "xn--ropox-9ved", "xn--drobx-xye4968b",
		"xn--dobox-0tb805b", "droplbox", "xn--drpdx-kyec",
		"xn--drpx-1ndb00h", "xn--drbox-kua12l", "xn--drpbx-mua201c",
		"xn--opbox-mwe55l", "d-ropbox", "dr-opbox",
		"dro-pbox", "drop-box", "dropb-ox",
		"dropbo-x", "droppbox", "drop0box",
		"dfropbox", "dr4opbox", "droipbox",
		"droopbox", "dreopbox", "dropbo9x",
		"dr9opbox", "dr5opbox", "drokpbox",
		"dtropbox", "drompbox", "dropbolx",
		"dropboix", "dropbo0x", "driopbox",
		"dropb0ox", "drpopbox", "drlopbox",
		"dro0pbox", "dropgbox", "dropbgox",
		"drtopbox", "deropbox", "dropmbox",
		"dropbiox", "dropblox", "dropbokx",
		"dropobox", "dropnbox", "drolpbox",
		"ddropbox", "drkopbox", "dropbnox",
		"drophbox", "dropb9ox", "d4ropbox",
		"dropbkox", "dropbopx", "dropbhox",
		"dropvbox", "drfopbox", "drdopbox",
		"d5ropbox", "dropbpox", "dro9pbox",
		"dr0opbox", "dropbvox", "dopbox",
		"dropbo", "drpbox", "dropbx",
		"drobox", "dropox", "ropbox",
		"dropboox", "dropbbox", "drropbox",
		"dr9pbox", "drlpbox", "dtopbox",
		"dropb0x", "dfopbox", "ddopbox",
		"d5opbox", "dr0pbox", "dropblx",
		"sropbox", "dropbod", "deopbox",
		"dropb9x", "dropboc", "xropbox",
		"d4opbox", "dropbpx", "dropgox",
		"dropbos", "dripbox", "droobox",
		"drppbox", "dropvox", "rropbox",
		"drolbox", "drombox", "dropnox",
		"cropbox", "dropbix", "drophox",
		"rdopbox", "dorpbox", "drpobox",
		"drobpox", "dropobx", "dropbxo",
		"dropbax", "dropbex", "drupbox",
		"dropbux", "drapbox", "drepbox",
		"dropboxcom",
	}

	return &Brand{
		Name:       name,
		Original:   original,
		Safelist:   safelist,
		Suspicious: suspicious,
	}
}

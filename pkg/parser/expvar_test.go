package parser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/wfernandes/app-metrics-plugin/pkg/parser"
)

var _ = Describe("Expvar", func() {

	It("removes expvar properties cmdline and memstats", func() {
		p := parser.NewExpvar([]string{"cmdline", "memstats"})
		output, err := p.Parse([]byte(expvarJSON))

		Expect(err).ToNot(HaveOccurred())
		Expect(output).To(MatchJSON(expectedJSON))
	})

	It("returns error when unable to unmarshal", func() {
		p := parser.NewExpvar([]string{"cmdline", "memstats"})
		_, err := p.Parse([]byte(`{"a": 123, "b": 456,}`))

		Expect(err).To(HaveOccurred())
	})

})

var expvarJSON = `{
"cmdline": ["bin/event-alerts"],
"ingress.matched": 11,
"ingress.received": 5664523,
"memstats": {"Alloc":1390288,"TotalAlloc":14425985608,"Sys":9509112,"Lookups":668,"Mallocs":126747222,"Frees":126735753,"HeapAlloc":1390288,"HeapSys":5505024,"HeapIdle":3293184,"HeapInuse":2211840,"HeapReleased":688128,"HeapObjects":11469,"StackInuse":786432,"StackSys":786432,"MSpanInuse":43320,"MSpanSys":98304,"MCacheInuse":4800,"MCacheSys":16384,"BuckHashSys":1459358,"GCSys":444416,"OtherSys":1199194,"NextGC":4194304,"LastGC":1511070581675115793,"PauseTotalNs":704230514,"PauseNs":[248046,80516,177301,195905,192759,155088,103042,201679,171202,101016,168150,235747,91611,98305,204128,260710,278451,80348,120232,101186,134314,176129,80528,100729,104065,157991,128525,97559,44751,122843,36948,154063,190383,150444,125052,175850,263287,182881,203957,180029,192082,85644,91101,161029,224597,108227,54009,167328,153760,186454,108018,151631,204805,100086,79952,177138,178176,121125,196399,153128,171934,93970,109306,190675,121725,148773,241194,165110,103737,196617,108381,77875,339701,103959,150459,132655,131317,219644,177542,281782,231313,143650,267856,166367,78417,244824,140053,74040,133315,185690,98041,156923,219383,114889,40856,144732,199442,128709,285413,154878,60979,190490,92276,99985,145981,168519,82811,135725,209615,178050,259315,58788,137927,170120,55973,170830,146744,167590,64546,115692,95334,134872,197791,154379,129539,147246,177524,78562,100456,231707,127347,123876,62460,142338,182237,138097,133739,227495,131776,89912,148908,128980,63346,112967,272607,184923,236354,111996,203516,67029,71648,184422,140741,107454,92071,210607,251378,267876,116741,231120,84114,111211,183277,152143,222563,56508,207041,179022,101331,177874,150134,181033,91086,175737,201887,127988,198702,176671,256045,182312,204702,132452,122707,116145,88330,69030,66989,142791,91264,1597333,87169,138373,77188,121883,128639,64754,188078,195048,213782,173951,220602,65439,197262,201025,117053,208684,120509,151241,237522,67969,147360,110897,83444,202017,254890,197441,241647,120181,210967,172153,221605,183217,120075,187864,257302,144890,92752,202938,76902,131740,118454,112095,146957,215135,165861,173173,72418,167797,113785,203961,175538,172851,236543,208609,168020,222869,116937,138989,201339,70893,124422,120276,180586,73777,155352,218545],"PauseEnd":[1511064986242940603,1511065010354812481,1511065036475156503,1511065061647734360,1511065085809811252,1511065111948897626,1511065136069365580,1511065163202281222,1511065187129126979,1511065213319117826,1511065237453962819,1511065262313379478,1511065289510039384,1511065313686167959,1511065339828745713,1511065363955037433,1511065390086039490,1511065415212953918,1511065439329921395,1511065466469399352,1511065490615564227,1511065516687675518,1511065540668226786,1511065565794607295,1511065591914448386,1511065617065014253,1511065643229665306,1511065667358085683,1511065693492678090,1511065718611348284,1511065742746152340,1511065769888891324,1511065793968357275,1511065821107355262,1511065843898082457,1511065871033259993,1511065895831279276,1511065919958495193,1511065947092555112,1511065971873052790,1511065998073978262,1511066022857565374,1511066047031149629,1511066074205039575,1511066098320098383,1511066124428688850,1511066148542722393,1511066175677306112,1511066199655691062,1511066224647674787,1511066250781696960,1511066275771752329,1511066301963054215,1511066325907579264,1511066351053678694,1511066375835859393,1511066400977637048,1511066428122119029,1511066451920014375,1511066478071766741,1511066503259380680,1511066527329224453,1511066554523435847,1511066578443721539,1511066605433937554,1511066630353052285,1511066656405847531,1511066680315704659,1511066704305812791,1511066730286288434,1511066755381455202,1511066781473848719,1511066806311343615,1511066832436268694,1511066857294512994,1511066881292072594,1511066908485975913,1511066933405159379,1511066960415148568,1511066984332958371,1511067010498898533,1511067035313991876,1511067060467353105,1511067086644416333,1511067110421139597,1511067136407858855,1511067161308721561,1511067185309638390,1511067212465355516,1511067237403439332,1511067263410155124,1511067287331433092,1511067314350345901,1511067338292156008,1511067363295653561,1511067390439354852,1511067415385673068,1511067440513879312,1511067465336523700,1511067490491180677,1511067515323392700,1511067540461119900,1511067567591828825,1511067591399613442,1511067617396079121,1511067642323922700,1511067666317203098,1511067693469116925,1511067718443515069,1511067744444174843,1511067768398676661,1511067794369921767,1511067819320442689,1511067843423335421,1511067870547236622,1511067895401026891,1511067920386519154,1511067945326788714,1511067970349125400,1511067996303499168,1511068020282904034,1511068046425250572,1511068071054167283,1511068097106680360,1511068121387400588,1511068146324245179,1511068171379586410,1511068196478646453,1511068223420739701,1511068247460211885,1511068273468586111,1511068297431890741,1511068321512358112,1511068349699744369,1511068373522882650,1511068400514582396,1511068424507183038,1511068450473997935,1511068475459952956,1511068500458472520,1511068526605795385,1511068550522260090,1511068576578873730,1511068601572185040,1511068625667805660,1511068651824270521,1511068676688094087,1511068702658293277,1511068727655118687,1511068753695175166,1511068777609529582,1511068802645853132,1511068828777405300,1511068853747916432,1511068879926092363,1511068903681752480,1511068929700868529,1511068955638925414,1511068978749762838,1511069005779766441,1511069030715594985,1511069055757552518,1511069081628905873,1511069105727915777,1511069130816202447,1511069155784541809,1511069180879030735,1511069205826916735,1511069231026312895,1511069256794957460,1511069280953889948,1511069308015345479,1511069331930608940,1511069358986359894,1511069382875193305,1511069408043876496,1511069435176805112,1511069458958820218,1511069485906549292,1511069509640685756,1511069535722875693,1511069560620992673,1511069584662132719,1511069609832783041,1511069635739085198,1511069660697102755,1511069685629164760,1511069710685663215,1511069735595698305,1511069760641553986,1511069786808069811,1511069811735648410,1511069836728903883,1511069857669605161,1511069879822906375,1511069905950938937,1511069927934768280,1511069952074647827,1511069976847107441,1511070001023980409,1511070027200332947,1511070051941525289,1511070079152255324,1511070103188307937,1511070128329911441,1511070155448597831,1511070178554212909,1511070205670773819,1511070230671389081,1511070255744295094,1511070280676855055,1511070304656274643,1511070330787717566,1511070355659766768,1511070381734320167,1511070405691422169,1511070431716010746,1511070456646754740,1511070480647253380,1511070507794847605,1511070531743648770,1511070557730033174,1511070581675115793,1511064155007621687,1511064180013031061,1511064206146901476,1511064230219843512,1511064256323620760,1511064280438855570,1511064305592882220,1511064330722245501,1511064356708943216,1511064382770577406,1511064406664258492,1511064432718140549,1511064455650182944,1511064480830172092,1511064507023524489,1511064531962814742,1511064557088949988,1511064582217899434,1511064606364515182,1511064633520262250,1511064657635444399,1511064683748629498,1511064707697896068,1511064733718423843,1511064757637934962,1511064781670019870,1511064808825874961,1511064832772592218,1511064858922520522,1511064883912385082,1511064909043517862,1511064934156059692,1511064959077610758],"NumGC":4319,"NumForcedGC":0,"GCCPUFraction":0.0000109818988389847,"EnableGC":true,"DebugGC":false,"BySize":[{"Size":0,"Mallocs":0,"Frees":0},{"Size":8,"Mallocs":5665721,"Frees":5665383},{"Size":16,"Mallocs":62669711,"Frees":62665684},{"Size":32,"Mallocs":3970944,"Frees":3966739},{"Size":48,"Mallocs":13241965,"Frees":13240977},{"Size":64,"Mallocs":83453,"Frees":83348},{"Size":80,"Mallocs":39586,"Frees":39316},{"Size":96,"Mallocs":8591,"Frees":8561},{"Size":112,"Mallocs":1104,"Frees":1013},{"Size":128,"Mallocs":632,"Frees":621},{"Size":144,"Mallocs":5672437,"Frees":5672114},{"Size":160,"Mallocs":26368,"Frees":26329},{"Size":176,"Mallocs":700,"Frees":674},{"Size":192,"Mallocs":130,"Frees":128},{"Size":208,"Mallocs":68,"Frees":45},{"Size":224,"Mallocs":317,"Frees":314},{"Size":240,"Mallocs":313,"Frees":311},{"Size":256,"Mallocs":364,"Frees":357},{"Size":288,"Mallocs":602144,"Frees":602091},{"Size":320,"Mallocs":96,"Frees":91},{"Size":352,"Mallocs":689,"Frees":677},{"Size":384,"Mallocs":91,"Frees":25},{"Size":416,"Mallocs":103,"Frees":6},{"Size":448,"Mallocs":15,"Frees":13},{"Size":480,"Mallocs":16,"Frees":13},{"Size":512,"Mallocs":5664579,"Frees":5664263},{"Size":576,"Mallocs":348,"Frees":344},{"Size":640,"Mallocs":29,"Frees":23},{"Size":704,"Mallocs":206,"Frees":197},{"Size":768,"Mallocs":16,"Frees":15},{"Size":896,"Mallocs":49,"Frees":42},{"Size":1024,"Mallocs":107,"Frees":92},{"Size":1152,"Mallocs":307,"Frees":305},{"Size":1280,"Mallocs":26,"Frees":20},{"Size":1408,"Mallocs":305,"Frees":302},{"Size":1536,"Mallocs":5664537,"Frees":5664221},{"Size":1792,"Mallocs":5,"Frees":1},{"Size":2048,"Mallocs":152,"Frees":133},{"Size":2304,"Mallocs":293,"Frees":287},{"Size":2688,"Mallocs":6,"Frees":0},{"Size":3072,"Mallocs":1,"Frees":1},{"Size":3200,"Mallocs":1,"Frees":0},{"Size":3456,"Mallocs":2,"Frees":0},{"Size":4096,"Mallocs":324,"Frees":320},{"Size":4864,"Mallocs":579,"Frees":574},{"Size":5376,"Mallocs":1,"Frees":0},{"Size":6144,"Mallocs":290,"Frees":288},{"Size":6528,"Mallocs":0,"Frees":0},{"Size":6784,"Mallocs":1,"Frees":0},{"Size":6912,"Mallocs":0,"Frees":0},{"Size":8192,"Mallocs":8,"Frees":7},{"Size":9472,"Mallocs":97,"Frees":97},{"Size":9728,"Mallocs":0,"Frees":0},{"Size":10240,"Mallocs":4,"Frees":4},{"Size":10880,"Mallocs":0,"Frees":0},{"Size":12288,"Mallocs":0,"Frees":0},{"Size":13568,"Mallocs":0,"Frees":0},{"Size":14336,"Mallocs":0,"Frees":0},{"Size":16384,"Mallocs":1,"Frees":0},{"Size":18432,"Mallocs":5,"Frees":4},{"Size":19072,"Mallocs":0,"Frees":0}]},
"notifier.dropped": 0,
"notifier.emails.failed": 0,
"notifier.emails.sent": 11
}`

var expectedJSON = `{
"ingress.matched": 11,
"ingress.received": 5664523,
"notifier.dropped": 0,
"notifier.emails.failed": 0,
"notifier.emails.sent": 11
}`

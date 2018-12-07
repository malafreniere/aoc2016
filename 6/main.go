package main

import (
	"fmt"
	"math"
	"strings"
)

const input = "vhrvabtc\nrzzdexux\npixjplcd\nimtxrwpe\njlowiwho\niqrfoytc\nulplfkix\nberacsou\nlnpbjpsd\ntjkoiwfm\nmbdwdtvc\nijzmhthl\nafuxnmmo\noalhgvyf\ncvnrvmmy\nphapcaaz\nqolkozza\ncslnokax\ncxmtylqr\ncelecybh\nlhsefcli\nncuttavx\ndwmgplhj\ndqclavvr\njgxgitqy\nlkermczl\npjaqasku\nnscotnwr\nlskrxypk\nrqufleiu\ndpebqpqm\nzrzqaknv\nejtyokhl\nzezxsymz\nvprfwlmo\nqbzpmtvp\nctrmxbgb\nvfuzasvt\neksijxxj\nrmehmkmw\nqtyvhcgk\nrlzkcewq\nvhnnobqv\nfxpeanog\nanldlmad\nirzcakuk\nbqwlsnlm\najqajujz\njdfncuag\nuogfnhhn\nkgovasxn\nahnehmpt\ndupgxqzj\nffvmbzcb\nhwwhqvuj\nftbirfob\nsfpupoze\ndcqqagnh\nlqcfxgui\nslaqsusj\nzfnlsbod\ngkkgjgoe\nbhhwadtt\nlkcahmbn\nbygxiuuq\nyfziyvix\ntgxfibtw\nkkmtqoko\ndnpradkn\nwhjuxqnf\nquxihhbp\nfvucbrjf\nseeczrga\nxxdtjtzm\ngibfcpip\npebojxbf\nsqwjhqau\nzcyoibrj\nvvrdfjlu\nnoiuyxer\ndqthjgoh\ntujhtems\ngoqyxish\nicyzsfvv\ntidfdydu\nwmkqxvhg\norojnwzl\nsoffzhen\nkcfpfeig\ndutcqtwp\nnskossbo\njdvatkfz\nlvamapuv\nolvjpcyz\nzkdaaszm\nxnutyfxf\njqfrlpdb\natjrmgbl\nsrmrnnez\ntahfxtwy\ncgmbsgtt\nmjidtwrh\nttaarqbz\necpeurzx\nbmperkew\nwrpkpzxy\njxqukdgp\npadrejjf\ncykpjmti\ngawyzbwk\nndwrqcnn\nqatvgyoo\nopptpjeb\nknfzqjsw\njwqvgpfu\nngvxvahq\nzcxpstzk\nscizdgxf\nrpijsqob\nfshiaxsc\nbkwyudab\nxndzhivu\nqabuelir\nxbbknsmi\naymcyblg\nudoozwuy\nbxwhsfnq\nrxfuclda\noviehgte\niwgomjxf\njxesodal\ntkkbgzvu\ncnbwowrp\nvsswuzxw\nqtgktieq\nnzgnvskg\nealykgra\nfeynkmtk\nlbkbojuq\nlawckzdr\nnserkqvj\nppnnvegc\nqgifeoix\nxztbecrg\ntiylyrsn\njfiramgc\nugezzfle\nmhsbmxqz\niaubkzne\nhozhztpo\nvujimqsj\njycbwqmv\nzeiuejjb\nbyuvzcyj\nikhpmoqb\nufjcsjgy\ncmcrbvst\nayunldxp\ngipqzdno\nqmsdgeal\nsycayzfs\nmszlvokk\nxqwossfj\njnqfghnm\npluhjlnk\nczlzxzbm\nkyywqhju\ndajubdhf\nkehffmfy\nmybcatox\nzctxuwwk\nybkxkawc\nzzwguxvv\nvlsswpcb\nvthsokmh\nzbwsbxws\nxstkrrwe\nhvsvgsgz\nqlqgzzmr\nhzuhclce\nyhfgpewt\neboyknhs\nnqjvjcqf\nnschhfdp\ngiupqixh\ndbzkbrsm\nulhllnde\nfzhmfhrr\nuiparxqh\ndxnnpfra\nludoygix\ndoqggixp\nomvihuem\nanijqdak\nhovtlipy\njvhulxyz\nazvkpvyc\nozutdhjc\nsjaaziki\nhghouxnn\ntuudbkez\ncfbjttzq\nsmtovjnc\nucipmtlj\ngxmatiyy\nkipdggkw\ngpnweeun\nfyzkpqbk\ndctgbfih\ntuhfntly\nxoanamuk\njlwxdjqp\nqzrnrjjz\nzffubqwg\nyujkbodi\nsrgfxrag\ntxvkzysv\needczuuz\ngosmfeex\neesqfxhg\nbjldoflf\nvvcbcnex\nwriyojbf\nutdjpovq\ndxrbgdff\nivwlxlfx\nuljthgnu\nasruzhqg\nqtxcvstv\nyfaylaqg\nwmnblvqo\nhuscfikc\neyqdkcvu\nnjbdiyhc\nhvmeacpb\nocxsnfca\neybztcgg\nrkqvbcst\numfkdhhp\nqepkcwqa\nxdnsscpw\nbsoaexkb\nkacsruck\nmwainqut\npdoqwphw\nuvblhueb\ncwrbiibu\nopcbzfgv\nzzyqzgrq\nqfgcrmpm\npeujlgqi\nnrfglgja\ntfnfbtbi\ncyqyqrtb\neqypdsvt\novdmqqvt\nhnukchfg\nbpfhlilo\nlkvwiusp\nzzhpkwbp\ngotqgblh\nmgcihbsc\ncvkatsjo\nnmtwcazc\nazenyubh\nujtoqfwh\nqdwkjcaf\nhurclzab\nmrbbwxxa\nrtosiexf\nvjfyfzol\nkwzltbnp\nakjpxpnd\nrclhgeyv\nudsvnvkd\ntrlzccds\nmdmcycpm\nyuqetosm\nyuyfnnbn\njwylebau\nybtynvca\npzsbkdxs\nipstuhyr\nrcnnyzrl\najgsynyt\nbzpifcnm\nebhvzfpm\nlaehwkbw\nfpinzejm\nhmevkjyz\npijsuvoz\nkpnjusky\nimhxxryz\nbbykcphd\neddgzrlt\nxofnilfp\npdclvqub\nxvqviwku\nahnximil\nzqmdysji\nsqoczvww\nxkmqsems\nqqraorxr\ndapbjpyi\nklospspq\nugqocfpx\nftjlpduh\nhpvzalgr\nlodrkkov\nxjtoerjr\nurfojzyy\njmbwqsdp\ntybiwbro\nnqguwpsg\nwzjdbtvp\nuamnromw\niiouocfq\nvbrifvdj\nzeylejin\niueydkui\nncyypjha\noalcyjvc\nhjtigsai\nyqbuzvod\nhzjxaoxm\nozdtjqrv\numltuvsn\ncyfxwaut\nfpnrkktw\nbsrtevsi\nmwkpttud\nlnmxrvnz\nohawlazx\npgwxgzva\nwiqpvalw\nlckdpqpj\nfzfadbcn\nmivmracr\nxlezdrfw\nescxeztx\ndehufqkb\nykaidzee\nkjxzlhxn\nwxearyjh\nfbtppgjb\njatslhys\nmxiwwyoj\nwjxsoisz\nqvlainlc\nwrpwzzds\nojylxjlk\nmnjgamga\nitdemmgx\neppmhvav\nzudnruzi\nfqhjdhrx\nrgsjwufu\ncfmcdsnl\nqmyludch\ndiczbwrc\nbptsolbs\ngregrfuk\nggtdtdsr\nvekxaxzq\niirnwtmj\njwimmqtl\nezrartxw\nmgzkopvx\nqtxhdjzl\nueoefpys\nwryfgobw\ngibrxmfq\nkuozniuj\nezctgdra\nfpxjfgww\nsoeicyjv\npookxsgw\nctatvacc\nymlpgowi\naekhewub\nbogedhje\naygzjewd\nbftmclcg\ntwspjgiu\nxflvkyzm\nweeihqco\nbzqissdv\nknhseano\npowynbnq\nylvewtrj\njwidlbzx\ndrienpvr\nrbndvvcn\nynhpvfds\nsgvvoexs\nynrsiahi\nyipvpmmm\nixxdnogd\nramwrfef\nndlumxte\ntvkqfyeo\nllkyyijo\ngfcamayi\nfdbicttv\ngiudovqr\njduryykk\nozsisrkr\nvymcyjnr\nfpoguwjh\nzzehbxvq\nbbzxtuco\ncreeinoi\npfgwqrcv\nljypthio\nesyytrqx\nvilobaon\nyharsrty\npcgygoak\nkdfzpikl\nolrosiss\nosivuyzn\novevojlo\ncwnxmqdo\npfmjddkq\nznvujmtf\nvkalccja\nnvzhfpia\nhyujfngl\nydefvfop\nkwiwltvo\nqjmowuox\nmenbophm\nzmljxiea\nicighnas\nnssgwyle\nwpxpiyjz\nvxnupnmu\nbbqsjnas\ntoveeodf\nootrclyu\njyklyedc\nrbktfwqd\nkbskmwua\nshnkkdta\nvxfkdxhh\ndebrhmid\nitrzmodt\nylaankij\nqcvlplpy\nxlwksxxk\nxdznurxi\nolwomzpf\nigqtqome\nxchlumws\nedcumpaz\npwjmdors\nlhxozxei\nhufvylqn\nedfevfzc\ngigdwfgf\nxggpedlq\nmlghgbfa\ngjacxeph\nsxvheyge\nckisqokq\nxndeieus\ngnxvdkwg\nahqogfbx\nnxzgxhep\nxpsarftf\nzmgrmeqr\nbhalyhol\nulejpxpy\nvhehrjzm\nwyqzpixv\nnkpjwbgq\nslwzvmua\ntubeisgw\nojgqkxdv\nkrjwyfcw\nmcittkjw\nsndxlapd\nxgdqjlcj\nrgozyuiz\nfprzkwln\ngwbgwqxm\nkhwmbdkn\nrwnqxwvk\nagxbmadn\nhyqnrdhw\ndbomlwal\nfbvmeuwt\nqsfgenhl\nudrfskjk\narajwwmn\nfghpeqbe\nbsjbfkdb\ngaixdesy\newchhofv\ntrjuuipw\neeumxzeh\nheqrqakk\nchjmlloy\nsmxcvpfi\nvxwonanp\npiygrizh\nsnxsnnfy\ntxsifcfx\nrqitketb\nmcbljita\nmwmgeogt\nsgvmpaph\nprmpddbr\nrxevhtdm\nlyejgafv\nmyveantx\nwoyyuycb\nsrovvzml\nvtudbulh\nxfouthsi\nkwgvwkzl\nmtwmvmxf\nddtwlrzr\ngtzwowtm\nhuzqtveo\ndcgmqyrn\ngtzzlhpe\nruvncyfs\nlwsxibps\nwaymjoru\nmmfdoxny\nrizwihlk\ncxaiotfr\nypoqbnry\nanlquyqk\nbhhffkkp\nkkqhromq\ntqowsiao\nfqkzblsd\ntoxdebej\nhozmnqou\nrcdonbwf\nusnojcyl\nufmpsswg\niefukrrp\nkuysztiv\nhnkytlvg\nnabjcyil\ntjtymamd\nzghngkdd\nmwbpuwmj\njlpqmghg\naslbbgqd\ndbutwlae\noqdimirz\nnpraupgz\ndakbvrfy\nyzrdqvag\nhrpaqgqj\nqakbxmpq\nutsimneu\nyjqtkmdz\nfqxlynyd\nwknsbcet\nkehwvrxu\nvxxrvhcd\nctpfdlsc\natjmqkye\npkaaqbvv\nfnmyewwe\ndtlqqbjd\nwsfyghha\nnndsvqbi\nwbaltpka\nzjuhspig\nisrerule\nqugfdfjt\nefjenzge\nwadqbjcc\ntvlcndmn\ncbveiint\nkkujfulr\nwtnyfzoe\nlmzgpulq\nihajkabg\npqiahufq\nravmkrdo\nldmftjct\nrhgdweob\nahjzgzkw\nvvgrueiq\nhiyxhnim\ngwynwuoc\nwonxvgee\njxwtxmau\nrggtslej\npbkwjyqf\nykanbgrg\nxqcwldyk\nnmlqdjws\nyvpcihki\nvwljaurh\nwucszart\nsshxhwid\nlfxnrqmt\njvvkxeht\nsiobukus\nedchngzp\nmkmuhoyd\nhidqtpid\nwvocqxbp\njvwqvvpt\ngrxtjncl\neyqiipyy\nyuxcubml\nfksvcuss\napdkwcqd\nbokchryh\nipbukstv\nifsynwcy\nodcgknnc\nomcsclgb\nzhjrucwj\nzloeydhr"
const example = "eedadn\ndrvtee\neandsr\nraavrd\natevrs\ntsrnev\nsdttsa\nrasrtv\nnssdts\nntnada\nsvetve\ntesnvt\nvntsnd\nvrdear\ndvrsen\nenarar"

type column map[rune]int
type columns [8]column

func main() {
	cols := columns{
		make(column),
		make(column),
		make(column),
		make(column),
		make(column),
		make(column),
		make(column),
		make(column),
	}

	for _, i := range strings.Split(input, "\n") {
		for j, c := range i {
			cols[j][c]++
		}
	}

	fmt.Println(cols.MaxString())
	fmt.Println(cols.MinString())
}

func (c *columns) MaxString() string {
	b := strings.Builder{}
	for _, col := range c {
		max := -1
		char := ' '

		for r, count := range col {
			if count > max {
				char = r
				max = count
			}
		}

		b.WriteRune(char)
	}
	return b.String()
}

func (c *columns) MinString() string {
	b := strings.Builder{}
	for _, col := range c {
		min := math.MaxInt32
		char := ' '

		for r, count := range col {
			if count < min {
				char = r
				min = count
			}
		}

		b.WriteRune(char)
	}
	return b.String()
}

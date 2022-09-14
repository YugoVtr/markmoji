package emoji

import "fmt"

type Emoji struct {
	IconURL string
	Code    int
}

func (e Emoji) String() string {
	return fmt.Sprintf("%s\t%s", string(e.Code), e.IconURL)
}

const BaseURL = "https://github.githubassets.com/images/icons/emoji/unicode"

type emojiIndexed map[string]Emoji

func (i emojiIndexed) String() (s string) {
	for name, code := range i {
		s = fmt.Sprintf("%s\n%-40s %s", s, name, code)
	}
	return fmt.Sprintln(s)
}

var Emojis = emojiIndexed{
	"+1": Emoji{
		IconURL: fmt.Sprintf("%s/1f44d.png?v8", BaseURL),
		Code:    128077,
	},
	"-1": Emoji{
		IconURL: fmt.Sprintf("%s/1f44e.png?v8", BaseURL),
		Code:    128078,
	},
	"100": Emoji{
		IconURL: fmt.Sprintf("%s/1f4af.png?v8", BaseURL),
		Code:    128175,
	},
	"1234": Emoji{
		IconURL: fmt.Sprintf("%s/1f522.png?v8", BaseURL),
		Code:    128290,
	},
	"1st_place_medal": Emoji{
		IconURL: fmt.Sprintf("%s/1f947.png?v8", BaseURL),
		Code:    129351,
	},
	"2nd_place_medal": Emoji{
		IconURL: fmt.Sprintf("%s/1f948.png?v8", BaseURL),
		Code:    129352,
	},
	"3rd_place_medal": Emoji{
		IconURL: fmt.Sprintf("%s/1f949.png?v8", BaseURL),
		Code:    129353,
	},
	"8ball": Emoji{
		IconURL: fmt.Sprintf("%s/1f3b1.png?v8", BaseURL),
		Code:    127921,
	},
	"a": Emoji{
		IconURL: fmt.Sprintf("%s/1f170.png?v8", BaseURL),
		Code:    127344,
	},
	"ab": Emoji{
		IconURL: fmt.Sprintf("%s/1f18e.png?v8", BaseURL),
		Code:    127374,
	},
	"abacus": Emoji{
		IconURL: fmt.Sprintf("%s/1f9ee.png?v8", BaseURL),
		Code:    129518,
	},
	"abc": Emoji{
		IconURL: fmt.Sprintf("%s/1f524.png?v8", BaseURL),
		Code:    128292,
	},
	"abcd": Emoji{
		IconURL: fmt.Sprintf("%s/1f521.png?v8", BaseURL),
		Code:    128289,
	},
	"accept": Emoji{
		IconURL: fmt.Sprintf("%s/1f251.png?v8", BaseURL),
		Code:    127569,
	},
	"accordion": Emoji{
		IconURL: fmt.Sprintf("%s/1fa97.png?v8", BaseURL),
		Code:    129687,
	},
	"adhesive_bandage": Emoji{
		IconURL: fmt.Sprintf("%s/1fa79.png?v8", BaseURL),
		Code:    129657,
	},
	"adult": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1.png?v8", BaseURL),
		Code:    129489,
	},
	"aerial_tramway": Emoji{
		IconURL: fmt.Sprintf("%s/1f6a1.png?v8", BaseURL),
		Code:    128673,
	},
	"afghanistan": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e6-1f1eb.png?v8", BaseURL),
		Code:    127462,
	},
	"airplane": Emoji{
		IconURL: fmt.Sprintf("%s/2708.png?v8", BaseURL),
		Code:    9992,
	},
	"aland_islands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e6-1f1fd.png?v8", BaseURL),
		Code:    127462,
	},
	"alarm_clock": Emoji{
		IconURL: fmt.Sprintf("%s/23f0.png?v8", BaseURL),
		Code:    9200,
	},
	"albania": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e6-1f1f1.png?v8", BaseURL),
		Code:    127462,
	},
	"alembic": Emoji{
		IconURL: fmt.Sprintf("%s/2697.png?v8", BaseURL),
		Code:    9879,
	},
	"algeria": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e9-1f1ff.png?v8", BaseURL),
		Code:    127465,
	},
	"alien": Emoji{
		IconURL: fmt.Sprintf("%s/1f47d.png?v8", BaseURL),
		Code:    128125,
	},
	"ambulance": Emoji{
		IconURL: fmt.Sprintf("%s/1f691.png?v8", BaseURL),
		Code:    128657,
	},
	"american_samoa": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e6-1f1f8.png?v8", BaseURL),
		Code:    127462,
	},
	"amphora": Emoji{
		IconURL: fmt.Sprintf("%s/1f3fa.png?v8", BaseURL),
		Code:    127994,
	},
	"anatomical_heart": Emoji{
		IconURL: fmt.Sprintf("%s/1fac0.png?v8", BaseURL),
		Code:    129728,
	},
	"anchor": Emoji{
		IconURL: fmt.Sprintf("%s/2693.png?v8", BaseURL),
		Code:    9875,
	},
	"andorra": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e6-1f1e9.png?v8", BaseURL),
		Code:    127462,
	},
	"angel": Emoji{
		IconURL: fmt.Sprintf("%s/1f47c.png?v8", BaseURL),
		Code:    128124,
	},
	"anger": Emoji{
		IconURL: fmt.Sprintf("%s/1f4a2.png?v8", BaseURL),
		Code:    128162,
	},
	"angola": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e6-1f1f4.png?v8", BaseURL),
		Code:    127462,
	},
	"angry": Emoji{
		IconURL: fmt.Sprintf("%s/1f620.png?v8", BaseURL),
		Code:    128544,
	},
	"anguilla": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e6-1f1ee.png?v8", BaseURL),
		Code:    127462,
	},
	"anguished": Emoji{
		IconURL: fmt.Sprintf("%s/1f627.png?v8", BaseURL),
		Code:    128551,
	},
	"ant": Emoji{
		IconURL: fmt.Sprintf("%s/1f41c.png?v8", BaseURL),
		Code:    128028,
	},
	"antarctica": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e6-1f1f6.png?v8", BaseURL),
		Code:    127462,
	},
	"antigua_barbuda": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e6-1f1ec.png?v8", BaseURL),
		Code:    127462,
	},
	"apple": Emoji{
		IconURL: fmt.Sprintf("%s/1f34e.png?v8", BaseURL),
		Code:    127822,
	},
	"aquarius": Emoji{
		IconURL: fmt.Sprintf("%s/2652.png?v8", BaseURL),
		Code:    9810,
	},
	"argentina": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e6-1f1f7.png?v8", BaseURL),
		Code:    127462,
	},
	"aries": Emoji{
		IconURL: fmt.Sprintf("%s/2648.png?v8", BaseURL),
		Code:    9800,
	},
	"armenia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e6-1f1f2.png?v8", BaseURL),
		Code:    127462,
	},
	"arrow_backward": Emoji{
		IconURL: fmt.Sprintf("%s/25c0.png?v8", BaseURL),
		Code:    9664,
	},
	"arrow_double_down": Emoji{
		IconURL: fmt.Sprintf("%s/23ec.png?v8", BaseURL),
		Code:    9196,
	},
	"arrow_double_up": Emoji{
		IconURL: fmt.Sprintf("%s/23eb.png?v8", BaseURL),
		Code:    9195,
	},
	"arrow_down": Emoji{
		IconURL: fmt.Sprintf("%s/2b07.png?v8", BaseURL),
		Code:    11015,
	},
	"arrow_down_small": Emoji{
		IconURL: fmt.Sprintf("%s/1f53d.png?v8", BaseURL),
		Code:    128317,
	},
	"arrow_forward": Emoji{
		IconURL: fmt.Sprintf("%s/25b6.png?v8", BaseURL),
		Code:    9654,
	},
	"arrow_heading_down": Emoji{
		IconURL: fmt.Sprintf("%s/2935.png?v8", BaseURL),
		Code:    10549,
	},
	"arrow_heading_up": Emoji{
		IconURL: fmt.Sprintf("%s/2934.png?v8", BaseURL),
		Code:    10548,
	},
	"arrow_left": Emoji{
		IconURL: fmt.Sprintf("%s/2b05.png?v8", BaseURL),
		Code:    11013,
	},
	"arrow_lower_left": Emoji{
		IconURL: fmt.Sprintf("%s/2199.png?v8", BaseURL),
		Code:    8601,
	},
	"arrow_lower_right": Emoji{
		IconURL: fmt.Sprintf("%s/2198.png?v8", BaseURL),
		Code:    8600,
	},
	"arrow_right": Emoji{
		IconURL: fmt.Sprintf("%s/27a1.png?v8", BaseURL),
		Code:    10145,
	},
	"arrow_right_hook": Emoji{
		IconURL: fmt.Sprintf("%s/21aa.png?v8", BaseURL),
		Code:    8618,
	},
	"arrow_up": Emoji{
		IconURL: fmt.Sprintf("%s/2b06.png?v8", BaseURL),
		Code:    11014,
	},
	"arrow_up_down": Emoji{
		IconURL: fmt.Sprintf("%s/2195.png?v8", BaseURL),
		Code:    8597,
	},
	"arrow_up_small": Emoji{
		IconURL: fmt.Sprintf("%s/1f53c.png?v8", BaseURL),
		Code:    128316,
	},
	"arrow_upper_left": Emoji{
		IconURL: fmt.Sprintf("%s/2196.png?v8", BaseURL),
		Code:    8598,
	},
	"arrow_upper_right": Emoji{
		IconURL: fmt.Sprintf("%s/2197.png?v8", BaseURL),
		Code:    8599,
	},
	"arrows_clockwise": Emoji{
		IconURL: fmt.Sprintf("%s/1f503.png?v8", BaseURL),
		Code:    128259,
	},
	"arrows_counterclockwise": Emoji{
		IconURL: fmt.Sprintf("%s/1f504.png?v8", BaseURL),
		Code:    128260,
	},
	"art": Emoji{
		IconURL: fmt.Sprintf("%s/1f3a8.png?v8", BaseURL),
		Code:    127912,
	},
	"articulated_lorry": Emoji{
		IconURL: fmt.Sprintf("%s/1f69b.png?v8", BaseURL),
		Code:    128667,
	},
	"artificial_satellite": Emoji{
		IconURL: fmt.Sprintf("%s/1f6f0.png?v8", BaseURL),
		Code:    128752,
	},
	"artist": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f3a8.png?v8", BaseURL),
		Code:    129489,
	},
	"aruba": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e6-1f1fc.png?v8", BaseURL),
		Code:    127462,
	},
	"ascension_island": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e6-1f1e8.png?v8", BaseURL),
		Code:    127462,
	},
	"asterisk": Emoji{
		IconURL: fmt.Sprintf("%s/002a-20e3.png?v8", BaseURL),
		Code:    42,
	},
	"astonished": Emoji{
		IconURL: fmt.Sprintf("%s/1f632.png?v8", BaseURL),
		Code:    128562,
	},
	"astronaut": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f680.png?v8", BaseURL),
		Code:    129489,
	},
	"athletic_shoe": Emoji{
		IconURL: fmt.Sprintf("%s/1f45f.png?v8", BaseURL),
		Code:    128095,
	},
	"atm": Emoji{
		IconURL: fmt.Sprintf("%s/1f3e7.png?v8", BaseURL),
		Code:    127975,
	},
	"atom": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/atom.png?v8",
		Code:    0,
	},
	"atom_symbol": Emoji{
		IconURL: fmt.Sprintf("%s/269b.png?v8", BaseURL),
		Code:    9883,
	},
	"australia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e6-1f1fa.png?v8", BaseURL),
		Code:    127462,
	},
	"austria": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e6-1f1f9.png?v8", BaseURL),
		Code:    127462,
	},
	"auto_rickshaw": Emoji{
		IconURL: fmt.Sprintf("%s/1f6fa.png?v8", BaseURL),
		Code:    128762,
	},
	"avocado": Emoji{
		IconURL: fmt.Sprintf("%s/1f951.png?v8", BaseURL),
		Code:    129361,
	},
	"axe": Emoji{
		IconURL: fmt.Sprintf("%s/1fa93.png?v8", BaseURL),
		Code:    129683,
	},
	"azerbaijan": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e6-1f1ff.png?v8", BaseURL),
		Code:    127462,
	},
	"b": Emoji{
		IconURL: fmt.Sprintf("%s/1f171.png?v8", BaseURL),
		Code:    127345,
	},
	"baby": Emoji{
		IconURL: fmt.Sprintf("%s/1f476.png?v8", BaseURL),
		Code:    128118,
	},
	"baby_bottle": Emoji{
		IconURL: fmt.Sprintf("%s/1f37c.png?v8", BaseURL),
		Code:    127868,
	},
	"baby_chick": Emoji{
		IconURL: fmt.Sprintf("%s/1f424.png?v8", BaseURL),
		Code:    128036,
	},
	"baby_symbol": Emoji{
		IconURL: fmt.Sprintf("%s/1f6bc.png?v8", BaseURL),
		Code:    128700,
	},
	"back": Emoji{
		IconURL: fmt.Sprintf("%s/1f519.png?v8", BaseURL),
		Code:    128281,
	},
	"bacon": Emoji{
		IconURL: fmt.Sprintf("%s/1f953.png?v8", BaseURL),
		Code:    129363,
	},
	"badger": Emoji{
		IconURL: fmt.Sprintf("%s/1f9a1.png?v8", BaseURL),
		Code:    129441,
	},
	"badminton": Emoji{
		IconURL: fmt.Sprintf("%s/1f3f8.png?v8", BaseURL),
		Code:    127992,
	},
	"bagel": Emoji{
		IconURL: fmt.Sprintf("%s/1f96f.png?v8", BaseURL),
		Code:    129391,
	},
	"baggage_claim": Emoji{
		IconURL: fmt.Sprintf("%s/1f6c4.png?v8", BaseURL),
		Code:    128708,
	},
	"baguette_bread": Emoji{
		IconURL: fmt.Sprintf("%s/1f956.png?v8", BaseURL),
		Code:    129366,
	},
	"bahamas": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1f8.png?v8", BaseURL),
		Code:    127463,
	},
	"bahrain": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1ed.png?v8", BaseURL),
		Code:    127463,
	},
	"balance_scale": Emoji{
		IconURL: fmt.Sprintf("%s/2696.png?v8", BaseURL),
		Code:    9878,
	},
	"bald_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f9b2.png?v8", BaseURL),
		Code:    128104,
	},
	"bald_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f9b2.png?v8", BaseURL),
		Code:    128105,
	},
	"ballet_shoes": Emoji{
		IconURL: fmt.Sprintf("%s/1fa70.png?v8", BaseURL),
		Code:    129648,
	},
	"balloon": Emoji{
		IconURL: fmt.Sprintf("%s/1f388.png?v8", BaseURL),
		Code:    127880,
	},
	"ballot_box": Emoji{
		IconURL: fmt.Sprintf("%s/1f5f3.png?v8", BaseURL),
		Code:    128499,
	},
	"ballot_box_with_check": Emoji{
		IconURL: fmt.Sprintf("%s/2611.png?v8", BaseURL),
		Code:    9745,
	},
	"bamboo": Emoji{
		IconURL: fmt.Sprintf("%s/1f38d.png?v8", BaseURL),
		Code:    127885,
	},
	"banana": Emoji{
		IconURL: fmt.Sprintf("%s/1f34c.png?v8", BaseURL),
		Code:    127820,
	},
	"bangbang": Emoji{
		IconURL: fmt.Sprintf("%s/203c.png?v8", BaseURL),
		Code:    8252,
	},
	"bangladesh": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1e9.png?v8", BaseURL),
		Code:    127463,
	},
	"banjo": Emoji{
		IconURL: fmt.Sprintf("%s/1fa95.png?v8", BaseURL),
		Code:    129685,
	},
	"bank": Emoji{
		IconURL: fmt.Sprintf("%s/1f3e6.png?v8", BaseURL),
		Code:    127974,
	},
	"bar_chart": Emoji{
		IconURL: fmt.Sprintf("%s/1f4ca.png?v8", BaseURL),
		Code:    128202,
	},
	"barbados": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1e7.png?v8", BaseURL),
		Code:    127463,
	},
	"barber": Emoji{
		IconURL: fmt.Sprintf("%s/1f488.png?v8", BaseURL),
		Code:    128136,
	},
	"baseball": Emoji{
		IconURL: fmt.Sprintf("%s/26be.png?v8", BaseURL),
		Code:    9918,
	},
	"basecamp": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/basecamp.png?v8",
		Code:    0,
	},
	"basecampy": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/basecampy.png?v8",
		Code:    0,
	},
	"basket": Emoji{
		IconURL: fmt.Sprintf("%s/1f9fa.png?v8", BaseURL),
		Code:    129530,
	},
	"basketball": Emoji{
		IconURL: fmt.Sprintf("%s/1f3c0.png?v8", BaseURL),
		Code:    127936,
	},
	"basketball_man": Emoji{
		IconURL: fmt.Sprintf("%s/26f9-2642.png?v8", BaseURL),
		Code:    9977,
	},
	"basketball_woman": Emoji{
		IconURL: fmt.Sprintf("%s/26f9-2640.png?v8", BaseURL),
		Code:    9977,
	},
	"bat": Emoji{
		IconURL: fmt.Sprintf("%s/1f987.png?v8", BaseURL),
		Code:    129415,
	},
	"bath": Emoji{
		IconURL: fmt.Sprintf("%s/1f6c0.png?v8", BaseURL),
		Code:    128704,
	},
	"bathtub": Emoji{
		IconURL: fmt.Sprintf("%s/1f6c1.png?v8", BaseURL),
		Code:    128705,
	},
	"battery": Emoji{
		IconURL: fmt.Sprintf("%s/1f50b.png?v8", BaseURL),
		Code:    128267,
	},
	"beach_umbrella": Emoji{
		IconURL: fmt.Sprintf("%s/1f3d6.png?v8", BaseURL),
		Code:    127958,
	},
	"bear": Emoji{
		IconURL: fmt.Sprintf("%s/1f43b.png?v8", BaseURL),
		Code:    128059,
	},
	"bearded_person": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d4.png?v8", BaseURL),
		Code:    129492,
	},
	"beaver": Emoji{
		IconURL: fmt.Sprintf("%s/1f9ab.png?v8", BaseURL),
		Code:    129451,
	},
	"bed": Emoji{
		IconURL: fmt.Sprintf("%s/1f6cf.png?v8", BaseURL),
		Code:    128719,
	},
	"bee": Emoji{
		IconURL: fmt.Sprintf("%s/1f41d.png?v8", BaseURL),
		Code:    128029,
	},
	"beer": Emoji{
		IconURL: fmt.Sprintf("%s/1f37a.png?v8", BaseURL),
		Code:    127866,
	},
	"beers": Emoji{
		IconURL: fmt.Sprintf("%s/1f37b.png?v8", BaseURL),
		Code:    127867,
	},
	"beetle": Emoji{
		IconURL: fmt.Sprintf("%s/1fab2.png?v8", BaseURL),
		Code:    129714,
	},
	"beginner": Emoji{
		IconURL: fmt.Sprintf("%s/1f530.png?v8", BaseURL),
		Code:    128304,
	},
	"belarus": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1fe.png?v8", BaseURL),
		Code:    127463,
	},
	"belgium": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1ea.png?v8", BaseURL),
		Code:    127463,
	},
	"belize": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1ff.png?v8", BaseURL),
		Code:    127463,
	},
	"bell": Emoji{
		IconURL: fmt.Sprintf("%s/1f514.png?v8", BaseURL),
		Code:    128276,
	},
	"bell_pepper": Emoji{
		IconURL: fmt.Sprintf("%s/1fad1.png?v8", BaseURL),
		Code:    129745,
	},
	"bellhop_bell": Emoji{
		IconURL: fmt.Sprintf("%s/1f6ce.png?v8", BaseURL),
		Code:    128718,
	},
	"benin": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1ef.png?v8", BaseURL),
		Code:    127463,
	},
	"bento": Emoji{
		IconURL: fmt.Sprintf("%s/1f371.png?v8", BaseURL),
		Code:    127857,
	},
	"bermuda": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1f2.png?v8", BaseURL),
		Code:    127463,
	},
	"beverage_box": Emoji{
		IconURL: fmt.Sprintf("%s/1f9c3.png?v8", BaseURL),
		Code:    129475,
	},
	"bhutan": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1f9.png?v8", BaseURL),
		Code:    127463,
	},
	"bicyclist": Emoji{
		IconURL: fmt.Sprintf("%s/1f6b4.png?v8", BaseURL),
		Code:    128692,
	},
	"bike": Emoji{
		IconURL: fmt.Sprintf("%s/1f6b2.png?v8", BaseURL),
		Code:    128690,
	},
	"biking_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f6b4-2642.png?v8", BaseURL),
		Code:    128692,
	},
	"biking_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f6b4-2640.png?v8", BaseURL),
		Code:    128692,
	},
	"bikini": Emoji{
		IconURL: fmt.Sprintf("%s/1f459.png?v8", BaseURL),
		Code:    128089,
	},
	"billed_cap": Emoji{
		IconURL: fmt.Sprintf("%s/1f9e2.png?v8", BaseURL),
		Code:    129506,
	},
	"biohazard": Emoji{
		IconURL: fmt.Sprintf("%s/2623.png?v8", BaseURL),
		Code:    9763,
	},
	"bird": Emoji{
		IconURL: fmt.Sprintf("%s/1f426.png?v8", BaseURL),
		Code:    128038,
	},
	"birthday": Emoji{
		IconURL: fmt.Sprintf("%s/1f382.png?v8", BaseURL),
		Code:    127874,
	},
	"bison": Emoji{
		IconURL: fmt.Sprintf("%s/1f9ac.png?v8", BaseURL),
		Code:    129452,
	},
	"black_cat": Emoji{
		IconURL: fmt.Sprintf("%s/1f408-2b1b.png?v8", BaseURL),
		Code:    128008,
	},
	"black_circle": Emoji{
		IconURL: fmt.Sprintf("%s/26ab.png?v8", BaseURL),
		Code:    9899,
	},
	"black_flag": Emoji{
		IconURL: fmt.Sprintf("%s/1f3f4.png?v8", BaseURL),
		Code:    127988,
	},
	"black_heart": Emoji{
		IconURL: fmt.Sprintf("%s/1f5a4.png?v8", BaseURL),
		Code:    128420,
	},
	"black_joker": Emoji{
		IconURL: fmt.Sprintf("%s/1f0cf.png?v8", BaseURL),
		Code:    127183,
	},
	"black_large_square": Emoji{
		IconURL: fmt.Sprintf("%s/2b1b.png?v8", BaseURL),
		Code:    11035,
	},
	"black_medium_small_square": Emoji{
		IconURL: fmt.Sprintf("%s/25fe.png?v8", BaseURL),
		Code:    9726,
	},
	"black_medium_square": Emoji{
		IconURL: fmt.Sprintf("%s/25fc.png?v8", BaseURL),
		Code:    9724,
	},
	"black_nib": Emoji{
		IconURL: fmt.Sprintf("%s/2712.png?v8", BaseURL),
		Code:    10002,
	},
	"black_small_square": Emoji{
		IconURL: fmt.Sprintf("%s/25aa.png?v8", BaseURL),
		Code:    9642,
	},
	"black_square_button": Emoji{
		IconURL: fmt.Sprintf("%s/1f532.png?v8", BaseURL),
		Code:    128306,
	},
	"blond_haired_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f471-2642.png?v8", BaseURL),
		Code:    128113,
	},
	"blond_haired_person": Emoji{
		IconURL: fmt.Sprintf("%s/1f471.png?v8", BaseURL),
		Code:    128113,
	},
	"blond_haired_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f471-2640.png?v8", BaseURL),
		Code:    128113,
	},
	"blonde_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f471-2640.png?v8", BaseURL),
		Code:    128113,
	},
	"blossom": Emoji{
		IconURL: fmt.Sprintf("%s/1f33c.png?v8", BaseURL),
		Code:    127804,
	},
	"blowfish": Emoji{
		IconURL: fmt.Sprintf("%s/1f421.png?v8", BaseURL),
		Code:    128033,
	},
	"blue_book": Emoji{
		IconURL: fmt.Sprintf("%s/1f4d8.png?v8", BaseURL),
		Code:    128216,
	},
	"blue_car": Emoji{
		IconURL: fmt.Sprintf("%s/1f699.png?v8", BaseURL),
		Code:    128665,
	},
	"blue_heart": Emoji{
		IconURL: fmt.Sprintf("%s/1f499.png?v8", BaseURL),
		Code:    128153,
	},
	"blue_square": Emoji{
		IconURL: fmt.Sprintf("%s/1f7e6.png?v8", BaseURL),
		Code:    128998,
	},
	"blueberries": Emoji{
		IconURL: fmt.Sprintf("%s/1fad0.png?v8", BaseURL),
		Code:    129744,
	},
	"blush": Emoji{
		IconURL: fmt.Sprintf("%s/1f60a.png?v8", BaseURL),
		Code:    128522,
	},
	"boar": Emoji{
		IconURL: fmt.Sprintf("%s/1f417.png?v8", BaseURL),
		Code:    128023,
	},
	"boat": Emoji{
		IconURL: fmt.Sprintf("%s/26f5.png?v8", BaseURL),
		Code:    9973,
	},
	"bolivia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1f4.png?v8", BaseURL),
		Code:    127463,
	},
	"bomb": Emoji{
		IconURL: fmt.Sprintf("%s/1f4a3.png?v8", BaseURL),
		Code:    128163,
	},
	"bone": Emoji{
		IconURL: fmt.Sprintf("%s/1f9b4.png?v8", BaseURL),
		Code:    129460,
	},
	"book": Emoji{
		IconURL: fmt.Sprintf("%s/1f4d6.png?v8", BaseURL),
		Code:    128214,
	},
	"bookmark": Emoji{
		IconURL: fmt.Sprintf("%s/1f516.png?v8", BaseURL),
		Code:    128278,
	},
	"bookmark_tabs": Emoji{
		IconURL: fmt.Sprintf("%s/1f4d1.png?v8", BaseURL),
		Code:    128209,
	},
	"books": Emoji{
		IconURL: fmt.Sprintf("%s/1f4da.png?v8", BaseURL),
		Code:    128218,
	},
	"boom": Emoji{
		IconURL: fmt.Sprintf("%s/1f4a5.png?v8", BaseURL),
		Code:    128165,
	},
	"boomerang": Emoji{
		IconURL: fmt.Sprintf("%s/1fa83.png?v8", BaseURL),
		Code:    129667,
	},
	"boot": Emoji{
		IconURL: fmt.Sprintf("%s/1f462.png?v8", BaseURL),
		Code:    128098,
	},
	"bosnia_herzegovina": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1e6.png?v8", BaseURL),
		Code:    127463,
	},
	"botswana": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1fc.png?v8", BaseURL),
		Code:    127463,
	},
	"bouncing_ball_man": Emoji{
		IconURL: fmt.Sprintf("%s/26f9-2642.png?v8", BaseURL),
		Code:    9977,
	},
	"bouncing_ball_person": Emoji{
		IconURL: fmt.Sprintf("%s/26f9.png?v8", BaseURL),
		Code:    9977,
	},
	"bouncing_ball_woman": Emoji{
		IconURL: fmt.Sprintf("%s/26f9-2640.png?v8", BaseURL),
		Code:    9977,
	},
	"bouquet": Emoji{
		IconURL: fmt.Sprintf("%s/1f490.png?v8", BaseURL),
		Code:    128144,
	},
	"bouvet_island": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1fb.png?v8", BaseURL),
		Code:    127463,
	},
	"bow": Emoji{
		IconURL: fmt.Sprintf("%s/1f647.png?v8", BaseURL),
		Code:    128583,
	},
	"bow_and_arrow": Emoji{
		IconURL: fmt.Sprintf("%s/1f3f9.png?v8", BaseURL),
		Code:    127993,
	},
	"bowing_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f647-2642.png?v8", BaseURL),
		Code:    128583,
	},
	"bowing_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f647-2640.png?v8", BaseURL),
		Code:    128583,
	},
	"bowl_with_spoon": Emoji{
		IconURL: fmt.Sprintf("%s/1f963.png?v8", BaseURL),
		Code:    129379,
	},
	"bowling": Emoji{
		IconURL: fmt.Sprintf("%s/1f3b3.png?v8", BaseURL),
		Code:    127923,
	},
	"bowtie": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/bowtie.png?v8",
		Code:    0,
	},
	"boxing_glove": Emoji{
		IconURL: fmt.Sprintf("%s/1f94a.png?v8", BaseURL),
		Code:    129354,
	},
	"boy": Emoji{
		IconURL: fmt.Sprintf("%s/1f466.png?v8", BaseURL),
		Code:    128102,
	},
	"brain": Emoji{
		IconURL: fmt.Sprintf("%s/1f9e0.png?v8", BaseURL),
		Code:    129504,
	},
	"brazil": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1f7.png?v8", BaseURL),
		Code:    127463,
	},
	"bread": Emoji{
		IconURL: fmt.Sprintf("%s/1f35e.png?v8", BaseURL),
		Code:    127838,
	},
	"breast_feeding": Emoji{
		IconURL: fmt.Sprintf("%s/1f931.png?v8", BaseURL),
		Code:    129329,
	},
	"bricks": Emoji{
		IconURL: fmt.Sprintf("%s/1f9f1.png?v8", BaseURL),
		Code:    129521,
	},
	"bride_with_veil": Emoji{
		IconURL: fmt.Sprintf("%s/1f470-2640.png?v8", BaseURL),
		Code:    128112,
	},
	"bridge_at_night": Emoji{
		IconURL: fmt.Sprintf("%s/1f309.png?v8", BaseURL),
		Code:    127753,
	},
	"briefcase": Emoji{
		IconURL: fmt.Sprintf("%s/1f4bc.png?v8", BaseURL),
		Code:    128188,
	},
	"british_indian_ocean_territory": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ee-1f1f4.png?v8", BaseURL),
		Code:    127470,
	},
	"british_virgin_islands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fb-1f1ec.png?v8", BaseURL),
		Code:    127483,
	},
	"broccoli": Emoji{
		IconURL: fmt.Sprintf("%s/1f966.png?v8", BaseURL),
		Code:    129382,
	},
	"broken_heart": Emoji{
		IconURL: fmt.Sprintf("%s/1f494.png?v8", BaseURL),
		Code:    128148,
	},
	"broom": Emoji{
		IconURL: fmt.Sprintf("%s/1f9f9.png?v8", BaseURL),
		Code:    129529,
	},
	"brown_circle": Emoji{
		IconURL: fmt.Sprintf("%s/1f7e4.png?v8", BaseURL),
		Code:    128996,
	},
	"brown_heart": Emoji{
		IconURL: fmt.Sprintf("%s/1f90e.png?v8", BaseURL),
		Code:    129294,
	},
	"brown_square": Emoji{
		IconURL: fmt.Sprintf("%s/1f7eb.png?v8", BaseURL),
		Code:    129003,
	},
	"brunei": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1f3.png?v8", BaseURL),
		Code:    127463,
	},
	"bubble_tea": Emoji{
		IconURL: fmt.Sprintf("%s/1f9cb.png?v8", BaseURL),
		Code:    129483,
	},
	"bucket": Emoji{
		IconURL: fmt.Sprintf("%s/1faa3.png?v8", BaseURL),
		Code:    129699,
	},
	"bug": Emoji{
		IconURL: fmt.Sprintf("%s/1f41b.png?v8", BaseURL),
		Code:    128027,
	},
	"building_construction": Emoji{
		IconURL: fmt.Sprintf("%s/1f3d7.png?v8", BaseURL),
		Code:    127959,
	},
	"bulb": Emoji{
		IconURL: fmt.Sprintf("%s/1f4a1.png?v8", BaseURL),
		Code:    128161,
	},
	"bulgaria": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1ec.png?v8", BaseURL),
		Code:    127463,
	},
	"bullettrain_front": Emoji{
		IconURL: fmt.Sprintf("%s/1f685.png?v8", BaseURL),
		Code:    128645,
	},
	"bullettrain_side": Emoji{
		IconURL: fmt.Sprintf("%s/1f684.png?v8", BaseURL),
		Code:    128644,
	},
	"burkina_faso": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1eb.png?v8", BaseURL),
		Code:    127463,
	},
	"burrito": Emoji{
		IconURL: fmt.Sprintf("%s/1f32f.png?v8", BaseURL),
		Code:    127791,
	},
	"burundi": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1ee.png?v8", BaseURL),
		Code:    127463,
	},
	"bus": Emoji{
		IconURL: fmt.Sprintf("%s/1f68c.png?v8", BaseURL),
		Code:    128652,
	},
	"business_suit_levitating": Emoji{
		IconURL: fmt.Sprintf("%s/1f574.png?v8", BaseURL),
		Code:    128372,
	},
	"busstop": Emoji{
		IconURL: fmt.Sprintf("%s/1f68f.png?v8", BaseURL),
		Code:    128655,
	},
	"bust_in_silhouette": Emoji{
		IconURL: fmt.Sprintf("%s/1f464.png?v8", BaseURL),
		Code:    128100,
	},
	"busts_in_silhouette": Emoji{
		IconURL: fmt.Sprintf("%s/1f465.png?v8", BaseURL),
		Code:    128101,
	},
	"butter": Emoji{
		IconURL: fmt.Sprintf("%s/1f9c8.png?v8", BaseURL),
		Code:    129480,
	},
	"butterfly": Emoji{
		IconURL: fmt.Sprintf("%s/1f98b.png?v8", BaseURL),
		Code:    129419,
	},
	"cactus": Emoji{
		IconURL: fmt.Sprintf("%s/1f335.png?v8", BaseURL),
		Code:    127797,
	},
	"cake": Emoji{
		IconURL: fmt.Sprintf("%s/1f370.png?v8", BaseURL),
		Code:    127856,
	},
	"calendar": Emoji{
		IconURL: fmt.Sprintf("%s/1f4c6.png?v8", BaseURL),
		Code:    128198,
	},
	"call_me_hand": Emoji{
		IconURL: fmt.Sprintf("%s/1f919.png?v8", BaseURL),
		Code:    129305,
	},
	"calling": Emoji{
		IconURL: fmt.Sprintf("%s/1f4f2.png?v8", BaseURL),
		Code:    128242,
	},
	"cambodia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f0-1f1ed.png?v8", BaseURL),
		Code:    127472,
	},
	"camel": Emoji{
		IconURL: fmt.Sprintf("%s/1f42b.png?v8", BaseURL),
		Code:    128043,
	},
	"camera": Emoji{
		IconURL: fmt.Sprintf("%s/1f4f7.png?v8", BaseURL),
		Code:    128247,
	},
	"camera_flash": Emoji{
		IconURL: fmt.Sprintf("%s/1f4f8.png?v8", BaseURL),
		Code:    128248,
	},
	"cameroon": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1f2.png?v8", BaseURL),
		Code:    127464,
	},
	"camping": Emoji{
		IconURL: fmt.Sprintf("%s/1f3d5.png?v8", BaseURL),
		Code:    127957,
	},
	"canada": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1e6.png?v8", BaseURL),
		Code:    127464,
	},
	"canary_islands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ee-1f1e8.png?v8", BaseURL),
		Code:    127470,
	},
	"cancer": Emoji{
		IconURL: fmt.Sprintf("%s/264b.png?v8", BaseURL),
		Code:    9803,
	},
	"candle": Emoji{
		IconURL: fmt.Sprintf("%s/1f56f.png?v8", BaseURL),
		Code:    128367,
	},
	"candy": Emoji{
		IconURL: fmt.Sprintf("%s/1f36c.png?v8", BaseURL),
		Code:    127852,
	},
	"canned_food": Emoji{
		IconURL: fmt.Sprintf("%s/1f96b.png?v8", BaseURL),
		Code:    129387,
	},
	"canoe": Emoji{
		IconURL: fmt.Sprintf("%s/1f6f6.png?v8", BaseURL),
		Code:    128758,
	},
	"cape_verde": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1fb.png?v8", BaseURL),
		Code:    127464,
	},
	"capital_abcd": Emoji{
		IconURL: fmt.Sprintf("%s/1f520.png?v8", BaseURL),
		Code:    128288,
	},
	"capricorn": Emoji{
		IconURL: fmt.Sprintf("%s/2651.png?v8", BaseURL),
		Code:    9809,
	},
	"car": Emoji{
		IconURL: fmt.Sprintf("%s/1f697.png?v8", BaseURL),
		Code:    128663,
	},
	"card_file_box": Emoji{
		IconURL: fmt.Sprintf("%s/1f5c3.png?v8", BaseURL),
		Code:    128451,
	},
	"card_index": Emoji{
		IconURL: fmt.Sprintf("%s/1f4c7.png?v8", BaseURL),
		Code:    128199,
	},
	"card_index_dividers": Emoji{
		IconURL: fmt.Sprintf("%s/1f5c2.png?v8", BaseURL),
		Code:    128450,
	},
	"caribbean_netherlands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1f6.png?v8", BaseURL),
		Code:    127463,
	},
	"carousel_horse": Emoji{
		IconURL: fmt.Sprintf("%s/1f3a0.png?v8", BaseURL),
		Code:    127904,
	},
	"carpentry_saw": Emoji{
		IconURL: fmt.Sprintf("%s/1fa9a.png?v8", BaseURL),
		Code:    129690,
	},
	"carrot": Emoji{
		IconURL: fmt.Sprintf("%s/1f955.png?v8", BaseURL),
		Code:    129365,
	},
	"cartwheeling": Emoji{
		IconURL: fmt.Sprintf("%s/1f938.png?v8", BaseURL),
		Code:    129336,
	},
	"cat": Emoji{
		IconURL: fmt.Sprintf("%s/1f431.png?v8", BaseURL),
		Code:    128049,
	},
	"cat2": Emoji{
		IconURL: fmt.Sprintf("%s/1f408.png?v8", BaseURL),
		Code:    128008,
	},
	"cayman_islands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f0-1f1fe.png?v8", BaseURL),
		Code:    127472,
	},
	"cd": Emoji{
		IconURL: fmt.Sprintf("%s/1f4bf.png?v8", BaseURL),
		Code:    128191,
	},
	"central_african_republic": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1eb.png?v8", BaseURL),
		Code:    127464,
	},
	"ceuta_melilla": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ea-1f1e6.png?v8", BaseURL),
		Code:    127466,
	},
	"chad": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f9-1f1e9.png?v8", BaseURL),
		Code:    127481,
	},
	"chains": Emoji{
		IconURL: fmt.Sprintf("%s/26d3.png?v8", BaseURL),
		Code:    9939,
	},
	"chair": Emoji{
		IconURL: fmt.Sprintf("%s/1fa91.png?v8", BaseURL),
		Code:    129681,
	},
	"champagne": Emoji{
		IconURL: fmt.Sprintf("%s/1f37e.png?v8", BaseURL),
		Code:    127870,
	},
	"chart": Emoji{
		IconURL: fmt.Sprintf("%s/1f4b9.png?v8", BaseURL),
		Code:    128185,
	},
	"chart_with_downwards_trend": Emoji{
		IconURL: fmt.Sprintf("%s/1f4c9.png?v8", BaseURL),
		Code:    128201,
	},
	"chart_with_upwards_trend": Emoji{
		IconURL: fmt.Sprintf("%s/1f4c8.png?v8", BaseURL),
		Code:    128200,
	},
	"checkered_flag": Emoji{
		IconURL: fmt.Sprintf("%s/1f3c1.png?v8", BaseURL),
		Code:    127937,
	},
	"cheese": Emoji{
		IconURL: fmt.Sprintf("%s/1f9c0.png?v8", BaseURL),
		Code:    129472,
	},
	"cherries": Emoji{
		IconURL: fmt.Sprintf("%s/1f352.png?v8", BaseURL),
		Code:    127826,
	},
	"cherry_blossom": Emoji{
		IconURL: fmt.Sprintf("%s/1f338.png?v8", BaseURL),
		Code:    127800,
	},
	"chess_pawn": Emoji{
		IconURL: fmt.Sprintf("%s/265f.png?v8", BaseURL),
		Code:    9823,
	},
	"chestnut": Emoji{
		IconURL: fmt.Sprintf("%s/1f330.png?v8", BaseURL),
		Code:    127792,
	},
	"chicken": Emoji{
		IconURL: fmt.Sprintf("%s/1f414.png?v8", BaseURL),
		Code:    128020,
	},
	"child": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d2.png?v8", BaseURL),
		Code:    129490,
	},
	"children_crossing": Emoji{
		IconURL: fmt.Sprintf("%s/1f6b8.png?v8", BaseURL),
		Code:    128696,
	},
	"chile": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1f1.png?v8", BaseURL),
		Code:    127464,
	},
	"chipmunk": Emoji{
		IconURL: fmt.Sprintf("%s/1f43f.png?v8", BaseURL),
		Code:    128063,
	},
	"chocolate_bar": Emoji{
		IconURL: fmt.Sprintf("%s/1f36b.png?v8", BaseURL),
		Code:    127851,
	},
	"chopsticks": Emoji{
		IconURL: fmt.Sprintf("%s/1f962.png?v8", BaseURL),
		Code:    129378,
	},
	"christmas_island": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1fd.png?v8", BaseURL),
		Code:    127464,
	},
	"christmas_tree": Emoji{
		IconURL: fmt.Sprintf("%s/1f384.png?v8", BaseURL),
		Code:    127876,
	},
	"church": Emoji{
		IconURL: fmt.Sprintf("%s/26ea.png?v8", BaseURL),
		Code:    9962,
	},
	"cinema": Emoji{
		IconURL: fmt.Sprintf("%s/1f3a6.png?v8", BaseURL),
		Code:    127910,
	},
	"circus_tent": Emoji{
		IconURL: fmt.Sprintf("%s/1f3aa.png?v8", BaseURL),
		Code:    127914,
	},
	"city_sunrise": Emoji{
		IconURL: fmt.Sprintf("%s/1f307.png?v8", BaseURL),
		Code:    127751,
	},
	"city_sunset": Emoji{
		IconURL: fmt.Sprintf("%s/1f306.png?v8", BaseURL),
		Code:    127750,
	},
	"cityscape": Emoji{
		IconURL: fmt.Sprintf("%s/1f3d9.png?v8", BaseURL),
		Code:    127961,
	},
	"cl": Emoji{
		IconURL: fmt.Sprintf("%s/1f191.png?v8", BaseURL),
		Code:    127377,
	},
	"clamp": Emoji{
		IconURL: fmt.Sprintf("%s/1f5dc.png?v8", BaseURL),
		Code:    128476,
	},
	"clap": Emoji{
		IconURL: fmt.Sprintf("%s/1f44f.png?v8", BaseURL),
		Code:    128079,
	},
	"clapper": Emoji{
		IconURL: fmt.Sprintf("%s/1f3ac.png?v8", BaseURL),
		Code:    127916,
	},
	"classical_building": Emoji{
		IconURL: fmt.Sprintf("%s/1f3db.png?v8", BaseURL),
		Code:    127963,
	},
	"climbing": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d7.png?v8", BaseURL),
		Code:    129495,
	},
	"climbing_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d7-2642.png?v8", BaseURL),
		Code:    129495,
	},
	"climbing_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d7-2640.png?v8", BaseURL),
		Code:    129495,
	},
	"clinking_glasses": Emoji{
		IconURL: fmt.Sprintf("%s/1f942.png?v8", BaseURL),
		Code:    129346,
	},
	"clipboard": Emoji{
		IconURL: fmt.Sprintf("%s/1f4cb.png?v8", BaseURL),
		Code:    128203,
	},
	"clipperton_island": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1f5.png?v8", BaseURL),
		Code:    127464,
	},
	"clock1": Emoji{
		IconURL: fmt.Sprintf("%s/1f550.png?v8", BaseURL),
		Code:    128336,
	},
	"clock10": Emoji{
		IconURL: fmt.Sprintf("%s/1f559.png?v8", BaseURL),
		Code:    128345,
	},
	"clock1030": Emoji{
		IconURL: fmt.Sprintf("%s/1f565.png?v8", BaseURL),
		Code:    128357,
	},
	"clock11": Emoji{
		IconURL: fmt.Sprintf("%s/1f55a.png?v8", BaseURL),
		Code:    128346,
	},
	"clock1130": Emoji{
		IconURL: fmt.Sprintf("%s/1f566.png?v8", BaseURL),
		Code:    128358,
	},
	"clock12": Emoji{
		IconURL: fmt.Sprintf("%s/1f55b.png?v8", BaseURL),
		Code:    128347,
	},
	"clock1230": Emoji{
		IconURL: fmt.Sprintf("%s/1f567.png?v8", BaseURL),
		Code:    128359,
	},
	"clock130": Emoji{
		IconURL: fmt.Sprintf("%s/1f55c.png?v8", BaseURL),
		Code:    128348,
	},
	"clock2": Emoji{
		IconURL: fmt.Sprintf("%s/1f551.png?v8", BaseURL),
		Code:    128337,
	},
	"clock230": Emoji{
		IconURL: fmt.Sprintf("%s/1f55d.png?v8", BaseURL),
		Code:    128349,
	},
	"clock3": Emoji{
		IconURL: fmt.Sprintf("%s/1f552.png?v8", BaseURL),
		Code:    128338,
	},
	"clock330": Emoji{
		IconURL: fmt.Sprintf("%s/1f55e.png?v8", BaseURL),
		Code:    128350,
	},
	"clock4": Emoji{
		IconURL: fmt.Sprintf("%s/1f553.png?v8", BaseURL),
		Code:    128339,
	},
	"clock430": Emoji{
		IconURL: fmt.Sprintf("%s/1f55f.png?v8", BaseURL),
		Code:    128351,
	},
	"clock5": Emoji{
		IconURL: fmt.Sprintf("%s/1f554.png?v8", BaseURL),
		Code:    128340,
	},
	"clock530": Emoji{
		IconURL: fmt.Sprintf("%s/1f560.png?v8", BaseURL),
		Code:    128352,
	},
	"clock6": Emoji{
		IconURL: fmt.Sprintf("%s/1f555.png?v8", BaseURL),
		Code:    128341,
	},
	"clock630": Emoji{
		IconURL: fmt.Sprintf("%s/1f561.png?v8", BaseURL),
		Code:    128353,
	},
	"clock7": Emoji{
		IconURL: fmt.Sprintf("%s/1f556.png?v8", BaseURL),
		Code:    128342,
	},
	"clock730": Emoji{
		IconURL: fmt.Sprintf("%s/1f562.png?v8", BaseURL),
		Code:    128354,
	},
	"clock8": Emoji{
		IconURL: fmt.Sprintf("%s/1f557.png?v8", BaseURL),
		Code:    128343,
	},
	"clock830": Emoji{
		IconURL: fmt.Sprintf("%s/1f563.png?v8", BaseURL),
		Code:    128355,
	},
	"clock9": Emoji{
		IconURL: fmt.Sprintf("%s/1f558.png?v8", BaseURL),
		Code:    128344,
	},
	"clock930": Emoji{
		IconURL: fmt.Sprintf("%s/1f564.png?v8", BaseURL),
		Code:    128356,
	},
	"closed_book": Emoji{
		IconURL: fmt.Sprintf("%s/1f4d5.png?v8", BaseURL),
		Code:    128213,
	},
	"closed_lock_with_key": Emoji{
		IconURL: fmt.Sprintf("%s/1f510.png?v8", BaseURL),
		Code:    128272,
	},
	"closed_umbrella": Emoji{
		IconURL: fmt.Sprintf("%s/1f302.png?v8", BaseURL),
		Code:    127746,
	},
	"cloud": Emoji{
		IconURL: fmt.Sprintf("%s/2601.png?v8", BaseURL),
		Code:    9729,
	},
	"cloud_with_lightning": Emoji{
		IconURL: fmt.Sprintf("%s/1f329.png?v8", BaseURL),
		Code:    127785,
	},
	"cloud_with_lightning_and_rain": Emoji{
		IconURL: fmt.Sprintf("%s/26c8.png?v8", BaseURL),
		Code:    9928,
	},
	"cloud_with_rain": Emoji{
		IconURL: fmt.Sprintf("%s/1f327.png?v8", BaseURL),
		Code:    127783,
	},
	"cloud_with_snow": Emoji{
		IconURL: fmt.Sprintf("%s/1f328.png?v8", BaseURL),
		Code:    127784,
	},
	"clown_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f921.png?v8", BaseURL),
		Code:    129313,
	},
	"clubs": Emoji{
		IconURL: fmt.Sprintf("%s/2663.png?v8", BaseURL),
		Code:    9827,
	},
	"cn": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1f3.png?v8", BaseURL),
		Code:    127464,
	},
	"coat": Emoji{
		IconURL: fmt.Sprintf("%s/1f9e5.png?v8", BaseURL),
		Code:    129509,
	},
	"cockroach": Emoji{
		IconURL: fmt.Sprintf("%s/1fab3.png?v8", BaseURL),
		Code:    129715,
	},
	"cocktail": Emoji{
		IconURL: fmt.Sprintf("%s/1f378.png?v8", BaseURL),
		Code:    127864,
	},
	"coconut": Emoji{
		IconURL: fmt.Sprintf("%s/1f965.png?v8", BaseURL),
		Code:    129381,
	},
	"cocos_islands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1e8.png?v8", BaseURL),
		Code:    127464,
	},
	"coffee": Emoji{
		IconURL: fmt.Sprintf("%s/2615.png?v8", BaseURL),
		Code:    9749,
	},
	"coffin": Emoji{
		IconURL: fmt.Sprintf("%s/26b0.png?v8", BaseURL),
		Code:    9904,
	},
	"coin": Emoji{
		IconURL: fmt.Sprintf("%s/1fa99.png?v8", BaseURL),
		Code:    129689,
	},
	"cold_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f976.png?v8", BaseURL),
		Code:    129398,
	},
	"cold_sweat": Emoji{
		IconURL: fmt.Sprintf("%s/1f630.png?v8", BaseURL),
		Code:    128560,
	},
	"collision": Emoji{
		IconURL: fmt.Sprintf("%s/1f4a5.png?v8", BaseURL),
		Code:    128165,
	},
	"colombia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1f4.png?v8", BaseURL),
		Code:    127464,
	},
	"comet": Emoji{
		IconURL: fmt.Sprintf("%s/2604.png?v8", BaseURL),
		Code:    9732,
	},
	"comoros": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f0-1f1f2.png?v8", BaseURL),
		Code:    127472,
	},
	"compass": Emoji{
		IconURL: fmt.Sprintf("%s/1f9ed.png?v8", BaseURL),
		Code:    129517,
	},
	"computer": Emoji{
		IconURL: fmt.Sprintf("%s/1f4bb.png?v8", BaseURL),
		Code:    128187,
	},
	"computer_mouse": Emoji{
		IconURL: fmt.Sprintf("%s/1f5b1.png?v8", BaseURL),
		Code:    128433,
	},
	"confetti_ball": Emoji{
		IconURL: fmt.Sprintf("%s/1f38a.png?v8", BaseURL),
		Code:    127882,
	},
	"confounded": Emoji{
		IconURL: fmt.Sprintf("%s/1f616.png?v8", BaseURL),
		Code:    128534,
	},
	"confused": Emoji{
		IconURL: fmt.Sprintf("%s/1f615.png?v8", BaseURL),
		Code:    128533,
	},
	"congo_brazzaville": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1ec.png?v8", BaseURL),
		Code:    127464,
	},
	"congo_kinshasa": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1e9.png?v8", BaseURL),
		Code:    127464,
	},
	"congratulations": Emoji{
		IconURL: fmt.Sprintf("%s/3297.png?v8", BaseURL),
		Code:    12951,
	},
	"construction": Emoji{
		IconURL: fmt.Sprintf("%s/1f6a7.png?v8", BaseURL),
		Code:    128679,
	},
	"construction_worker": Emoji{
		IconURL: fmt.Sprintf("%s/1f477.png?v8", BaseURL),
		Code:    128119,
	},
	"construction_worker_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f477-2642.png?v8", BaseURL),
		Code:    128119,
	},
	"construction_worker_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f477-2640.png?v8", BaseURL),
		Code:    128119,
	},
	"control_knobs": Emoji{
		IconURL: fmt.Sprintf("%s/1f39b.png?v8", BaseURL),
		Code:    127899,
	},
	"convenience_store": Emoji{
		IconURL: fmt.Sprintf("%s/1f3ea.png?v8", BaseURL),
		Code:    127978,
	},
	"cook": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f373.png?v8", BaseURL),
		Code:    129489,
	},
	"cook_islands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1f0.png?v8", BaseURL),
		Code:    127464,
	},
	"cookie": Emoji{
		IconURL: fmt.Sprintf("%s/1f36a.png?v8", BaseURL),
		Code:    127850,
	},
	"cool": Emoji{
		IconURL: fmt.Sprintf("%s/1f192.png?v8", BaseURL),
		Code:    127378,
	},
	"cop": Emoji{
		IconURL: fmt.Sprintf("%s/1f46e.png?v8", BaseURL),
		Code:    128110,
	},
	"copyright": Emoji{
		IconURL: fmt.Sprintf("%s/00a9.png?v8", BaseURL),
		Code:    169,
	},
	"corn": Emoji{
		IconURL: fmt.Sprintf("%s/1f33d.png?v8", BaseURL),
		Code:    127805,
	},
	"costa_rica": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1f7.png?v8", BaseURL),
		Code:    127464,
	},
	"cote_divoire": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1ee.png?v8", BaseURL),
		Code:    127464,
	},
	"couch_and_lamp": Emoji{
		IconURL: fmt.Sprintf("%s/1f6cb.png?v8", BaseURL),
		Code:    128715,
	},
	"couple": Emoji{
		IconURL: fmt.Sprintf("%s/1f46b.png?v8", BaseURL),
		Code:    128107,
	},
	"couple_with_heart": Emoji{
		IconURL: fmt.Sprintf("%s/1f491.png?v8", BaseURL),
		Code:    128145,
	},
	"couple_with_heart_man_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-2764-1f468.png?v8", BaseURL),
		Code:    128104,
	},
	"couple_with_heart_woman_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-2764-1f468.png?v8", BaseURL),
		Code:    128105,
	},
	"couple_with_heart_woman_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-2764-1f469.png?v8", BaseURL),
		Code:    128105,
	},
	"couplekiss": Emoji{
		IconURL: fmt.Sprintf("%s/1f48f.png?v8", BaseURL),
		Code:    128143,
	},
	"couplekiss_man_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-2764-1f48b-1f468.png?v8", BaseURL),
		Code:    128104,
	},
	"couplekiss_man_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-2764-1f48b-1f468.png?v8", BaseURL),
		Code:    128105,
	},
	"couplekiss_woman_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-2764-1f48b-1f469.png?v8", BaseURL),
		Code:    128105,
	},
	"cow": Emoji{
		IconURL: fmt.Sprintf("%s/1f42e.png?v8", BaseURL),
		Code:    128046,
	},
	"cow2": Emoji{
		IconURL: fmt.Sprintf("%s/1f404.png?v8", BaseURL),
		Code:    128004,
	},
	"cowboy_hat_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f920.png?v8", BaseURL),
		Code:    129312,
	},
	"crab": Emoji{
		IconURL: fmt.Sprintf("%s/1f980.png?v8", BaseURL),
		Code:    129408,
	},
	"crayon": Emoji{
		IconURL: fmt.Sprintf("%s/1f58d.png?v8", BaseURL),
		Code:    128397,
	},
	"credit_card": Emoji{
		IconURL: fmt.Sprintf("%s/1f4b3.png?v8", BaseURL),
		Code:    128179,
	},
	"crescent_moon": Emoji{
		IconURL: fmt.Sprintf("%s/1f319.png?v8", BaseURL),
		Code:    127769,
	},
	"cricket": Emoji{
		IconURL: fmt.Sprintf("%s/1f997.png?v8", BaseURL),
		Code:    129431,
	},
	"cricket_game": Emoji{
		IconURL: fmt.Sprintf("%s/1f3cf.png?v8", BaseURL),
		Code:    127951,
	},
	"croatia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ed-1f1f7.png?v8", BaseURL),
		Code:    127469,
	},
	"crocodile": Emoji{
		IconURL: fmt.Sprintf("%s/1f40a.png?v8", BaseURL),
		Code:    128010,
	},
	"croissant": Emoji{
		IconURL: fmt.Sprintf("%s/1f950.png?v8", BaseURL),
		Code:    129360,
	},
	"crossed_fingers": Emoji{
		IconURL: fmt.Sprintf("%s/1f91e.png?v8", BaseURL),
		Code:    129310,
	},
	"crossed_flags": Emoji{
		IconURL: fmt.Sprintf("%s/1f38c.png?v8", BaseURL),
		Code:    127884,
	},
	"crossed_swords": Emoji{
		IconURL: fmt.Sprintf("%s/2694.png?v8", BaseURL),
		Code:    9876,
	},
	"crown": Emoji{
		IconURL: fmt.Sprintf("%s/1f451.png?v8", BaseURL),
		Code:    128081,
	},
	"cry": Emoji{
		IconURL: fmt.Sprintf("%s/1f622.png?v8", BaseURL),
		Code:    128546,
	},
	"crying_cat_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f63f.png?v8", BaseURL),
		Code:    128575,
	},
	"crystal_ball": Emoji{
		IconURL: fmt.Sprintf("%s/1f52e.png?v8", BaseURL),
		Code:    128302,
	},
	"cuba": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1fa.png?v8", BaseURL),
		Code:    127464,
	},
	"cucumber": Emoji{
		IconURL: fmt.Sprintf("%s/1f952.png?v8", BaseURL),
		Code:    129362,
	},
	"cup_with_straw": Emoji{
		IconURL: fmt.Sprintf("%s/1f964.png?v8", BaseURL),
		Code:    129380,
	},
	"cupcake": Emoji{
		IconURL: fmt.Sprintf("%s/1f9c1.png?v8", BaseURL),
		Code:    129473,
	},
	"cupid": Emoji{
		IconURL: fmt.Sprintf("%s/1f498.png?v8", BaseURL),
		Code:    128152,
	},
	"curacao": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1fc.png?v8", BaseURL),
		Code:    127464,
	},
	"curling_stone": Emoji{
		IconURL: fmt.Sprintf("%s/1f94c.png?v8", BaseURL),
		Code:    129356,
	},
	"curly_haired_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f9b1.png?v8", BaseURL),
		Code:    128104,
	},
	"curly_haired_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f9b1.png?v8", BaseURL),
		Code:    128105,
	},
	"curly_loop": Emoji{
		IconURL: fmt.Sprintf("%s/27b0.png?v8", BaseURL),
		Code:    10160,
	},
	"currency_exchange": Emoji{
		IconURL: fmt.Sprintf("%s/1f4b1.png?v8", BaseURL),
		Code:    128177,
	},
	"curry": Emoji{
		IconURL: fmt.Sprintf("%s/1f35b.png?v8", BaseURL),
		Code:    127835,
	},
	"cursing_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f92c.png?v8", BaseURL),
		Code:    129324,
	},
	"custard": Emoji{
		IconURL: fmt.Sprintf("%s/1f36e.png?v8", BaseURL),
		Code:    127854,
	},
	"customs": Emoji{
		IconURL: fmt.Sprintf("%s/1f6c3.png?v8", BaseURL),
		Code:    128707,
	},
	"cut_of_meat": Emoji{
		IconURL: fmt.Sprintf("%s/1f969.png?v8", BaseURL),
		Code:    129385,
	},
	"cyclone": Emoji{
		IconURL: fmt.Sprintf("%s/1f300.png?v8", BaseURL),
		Code:    127744,
	},
	"cyprus": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1fe.png?v8", BaseURL),
		Code:    127464,
	},
	"czech_republic": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1ff.png?v8", BaseURL),
		Code:    127464,
	},
	"dagger": Emoji{
		IconURL: fmt.Sprintf("%s/1f5e1.png?v8", BaseURL),
		Code:    128481,
	},
	"dancer": Emoji{
		IconURL: fmt.Sprintf("%s/1f483.png?v8", BaseURL),
		Code:    128131,
	},
	"dancers": Emoji{
		IconURL: fmt.Sprintf("%s/1f46f.png?v8", BaseURL),
		Code:    128111,
	},
	"dancing_men": Emoji{
		IconURL: fmt.Sprintf("%s/1f46f-2642.png?v8", BaseURL),
		Code:    128111,
	},
	"dancing_women": Emoji{
		IconURL: fmt.Sprintf("%s/1f46f-2640.png?v8", BaseURL),
		Code:    128111,
	},
	"dango": Emoji{
		IconURL: fmt.Sprintf("%s/1f361.png?v8", BaseURL),
		Code:    127841,
	},
	"dark_sunglasses": Emoji{
		IconURL: fmt.Sprintf("%s/1f576.png?v8", BaseURL),
		Code:    128374,
	},
	"dart": Emoji{
		IconURL: fmt.Sprintf("%s/1f3af.png?v8", BaseURL),
		Code:    127919,
	},
	"dash": Emoji{
		IconURL: fmt.Sprintf("%s/1f4a8.png?v8", BaseURL),
		Code:    128168,
	},
	"date": Emoji{
		IconURL: fmt.Sprintf("%s/1f4c5.png?v8", BaseURL),
		Code:    128197,
	},
	"de": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e9-1f1ea.png?v8", BaseURL),
		Code:    127465,
	},
	"deaf_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f9cf-2642.png?v8", BaseURL),
		Code:    129487,
	},
	"deaf_person": Emoji{
		IconURL: fmt.Sprintf("%s/1f9cf.png?v8", BaseURL),
		Code:    129487,
	},
	"deaf_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f9cf-2640.png?v8", BaseURL),
		Code:    129487,
	},
	"deciduous_tree": Emoji{
		IconURL: fmt.Sprintf("%s/1f333.png?v8", BaseURL),
		Code:    127795,
	},
	"deer": Emoji{
		IconURL: fmt.Sprintf("%s/1f98c.png?v8", BaseURL),
		Code:    129420,
	},
	"denmark": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e9-1f1f0.png?v8", BaseURL),
		Code:    127465,
	},
	"department_store": Emoji{
		IconURL: fmt.Sprintf("%s/1f3ec.png?v8", BaseURL),
		Code:    127980,
	},
	"derelict_house": Emoji{
		IconURL: fmt.Sprintf("%s/1f3da.png?v8", BaseURL),
		Code:    127962,
	},
	"desert": Emoji{
		IconURL: fmt.Sprintf("%s/1f3dc.png?v8", BaseURL),
		Code:    127964,
	},
	"desert_island": Emoji{
		IconURL: fmt.Sprintf("%s/1f3dd.png?v8", BaseURL),
		Code:    127965,
	},
	"desktop_computer": Emoji{
		IconURL: fmt.Sprintf("%s/1f5a5.png?v8", BaseURL),
		Code:    128421,
	},
	"detective": Emoji{
		IconURL: fmt.Sprintf("%s/1f575.png?v8", BaseURL),
		Code:    128373,
	},
	"diamond_shape_with_a_dot_inside": Emoji{
		IconURL: fmt.Sprintf("%s/1f4a0.png?v8", BaseURL),
		Code:    128160,
	},
	"diamonds": Emoji{
		IconURL: fmt.Sprintf("%s/2666.png?v8", BaseURL),
		Code:    9830,
	},
	"diego_garcia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e9-1f1ec.png?v8", BaseURL),
		Code:    127465,
	},
	"disappointed": Emoji{
		IconURL: fmt.Sprintf("%s/1f61e.png?v8", BaseURL),
		Code:    128542,
	},
	"disappointed_relieved": Emoji{
		IconURL: fmt.Sprintf("%s/1f625.png?v8", BaseURL),
		Code:    128549,
	},
	"disguised_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f978.png?v8", BaseURL),
		Code:    129400,
	},
	"diving_mask": Emoji{
		IconURL: fmt.Sprintf("%s/1f93f.png?v8", BaseURL),
		Code:    129343,
	},
	"diya_lamp": Emoji{
		IconURL: fmt.Sprintf("%s/1fa94.png?v8", BaseURL),
		Code:    129684,
	},
	"dizzy": Emoji{
		IconURL: fmt.Sprintf("%s/1f4ab.png?v8", BaseURL),
		Code:    128171,
	},
	"dizzy_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f635.png?v8", BaseURL),
		Code:    128565,
	},
	"djibouti": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e9-1f1ef.png?v8", BaseURL),
		Code:    127465,
	},
	"dna": Emoji{
		IconURL: fmt.Sprintf("%s/1f9ec.png?v8", BaseURL),
		Code:    129516,
	},
	"do_not_litter": Emoji{
		IconURL: fmt.Sprintf("%s/1f6af.png?v8", BaseURL),
		Code:    128687,
	},
	"dodo": Emoji{
		IconURL: fmt.Sprintf("%s/1f9a4.png?v8", BaseURL),
		Code:    129444,
	},
	"dog": Emoji{
		IconURL: fmt.Sprintf("%s/1f436.png?v8", BaseURL),
		Code:    128054,
	},
	"dog2": Emoji{
		IconURL: fmt.Sprintf("%s/1f415.png?v8", BaseURL),
		Code:    128021,
	},
	"dollar": Emoji{
		IconURL: fmt.Sprintf("%s/1f4b5.png?v8", BaseURL),
		Code:    128181,
	},
	"dolls": Emoji{
		IconURL: fmt.Sprintf("%s/1f38e.png?v8", BaseURL),
		Code:    127886,
	},
	"dolphin": Emoji{
		IconURL: fmt.Sprintf("%s/1f42c.png?v8", BaseURL),
		Code:    128044,
	},
	"dominica": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e9-1f1f2.png?v8", BaseURL),
		Code:    127465,
	},
	"dominican_republic": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e9-1f1f4.png?v8", BaseURL),
		Code:    127465,
	},
	"door": Emoji{
		IconURL: fmt.Sprintf("%s/1f6aa.png?v8", BaseURL),
		Code:    128682,
	},
	"doughnut": Emoji{
		IconURL: fmt.Sprintf("%s/1f369.png?v8", BaseURL),
		Code:    127849,
	},
	"dove": Emoji{
		IconURL: fmt.Sprintf("%s/1f54a.png?v8", BaseURL),
		Code:    128330,
	},
	"dragon": Emoji{
		IconURL: fmt.Sprintf("%s/1f409.png?v8", BaseURL),
		Code:    128009,
	},
	"dragon_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f432.png?v8", BaseURL),
		Code:    128050,
	},
	"dress": Emoji{
		IconURL: fmt.Sprintf("%s/1f457.png?v8", BaseURL),
		Code:    128087,
	},
	"dromedary_camel": Emoji{
		IconURL: fmt.Sprintf("%s/1f42a.png?v8", BaseURL),
		Code:    128042,
	},
	"drooling_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f924.png?v8", BaseURL),
		Code:    129316,
	},
	"drop_of_blood": Emoji{
		IconURL: fmt.Sprintf("%s/1fa78.png?v8", BaseURL),
		Code:    129656,
	},
	"droplet": Emoji{
		IconURL: fmt.Sprintf("%s/1f4a7.png?v8", BaseURL),
		Code:    128167,
	},
	"drum": Emoji{
		IconURL: fmt.Sprintf("%s/1f941.png?v8", BaseURL),
		Code:    129345,
	},
	"duck": Emoji{
		IconURL: fmt.Sprintf("%s/1f986.png?v8", BaseURL),
		Code:    129414,
	},
	"dumpling": Emoji{
		IconURL: fmt.Sprintf("%s/1f95f.png?v8", BaseURL),
		Code:    129375,
	},
	"dvd": Emoji{
		IconURL: fmt.Sprintf("%s/1f4c0.png?v8", BaseURL),
		Code:    128192,
	},
	"e-mail": Emoji{
		IconURL: fmt.Sprintf("%s/1f4e7.png?v8", BaseURL),
		Code:    128231,
	},
	"eagle": Emoji{
		IconURL: fmt.Sprintf("%s/1f985.png?v8", BaseURL),
		Code:    129413,
	},
	"ear": Emoji{
		IconURL: fmt.Sprintf("%s/1f442.png?v8", BaseURL),
		Code:    128066,
	},
	"ear_of_rice": Emoji{
		IconURL: fmt.Sprintf("%s/1f33e.png?v8", BaseURL),
		Code:    127806,
	},
	"ear_with_hearing_aid": Emoji{
		IconURL: fmt.Sprintf("%s/1f9bb.png?v8", BaseURL),
		Code:    129467,
	},
	"earth_africa": Emoji{
		IconURL: fmt.Sprintf("%s/1f30d.png?v8", BaseURL),
		Code:    127757,
	},
	"earth_americas": Emoji{
		IconURL: fmt.Sprintf("%s/1f30e.png?v8", BaseURL),
		Code:    127758,
	},
	"earth_asia": Emoji{
		IconURL: fmt.Sprintf("%s/1f30f.png?v8", BaseURL),
		Code:    127759,
	},
	"ecuador": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ea-1f1e8.png?v8", BaseURL),
		Code:    127466,
	},
	"egg": Emoji{
		IconURL: fmt.Sprintf("%s/1f95a.png?v8", BaseURL),
		Code:    129370,
	},
	"eggplant": Emoji{
		IconURL: fmt.Sprintf("%s/1f346.png?v8", BaseURL),
		Code:    127814,
	},
	"egypt": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ea-1f1ec.png?v8", BaseURL),
		Code:    127466,
	},
	"eight": Emoji{
		IconURL: fmt.Sprintf("%s/0038-20e3.png?v8", BaseURL),
		Code:    56,
	},
	"eight_pointed_black_star": Emoji{
		IconURL: fmt.Sprintf("%s/2734.png?v8", BaseURL),
		Code:    10036,
	},
	"eight_spoked_asterisk": Emoji{
		IconURL: fmt.Sprintf("%s/2733.png?v8", BaseURL),
		Code:    10035,
	},
	"eject_button": Emoji{
		IconURL: fmt.Sprintf("%s/23cf.png?v8", BaseURL),
		Code:    9167,
	},
	"el_salvador": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1fb.png?v8", BaseURL),
		Code:    127480,
	},
	"electric_plug": Emoji{
		IconURL: fmt.Sprintf("%s/1f50c.png?v8", BaseURL),
		Code:    128268,
	},
	"electron": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/electron.png?v8",
		Code:    0,
	},
	"elephant": Emoji{
		IconURL: fmt.Sprintf("%s/1f418.png?v8", BaseURL),
		Code:    128024,
	},
	"elevator": Emoji{
		IconURL: fmt.Sprintf("%s/1f6d7.png?v8", BaseURL),
		Code:    128727,
	},
	"elf": Emoji{
		IconURL: fmt.Sprintf("%s/1f9dd.png?v8", BaseURL),
		Code:    129501,
	},
	"elf_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f9dd-2642.png?v8", BaseURL),
		Code:    129501,
	},
	"elf_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f9dd-2640.png?v8", BaseURL),
		Code:    129501,
	},
	"email": Emoji{
		IconURL: fmt.Sprintf("%s/1f4e7.png?v8", BaseURL),
		Code:    128231,
	},
	"end": Emoji{
		IconURL: fmt.Sprintf("%s/1f51a.png?v8", BaseURL),
		Code:    128282,
	},
	"england": Emoji{
		IconURL: fmt.Sprintf("%s/1f3f4-e0067-e0062-e0065-e006e-e0067-e007f.png?v8", BaseURL),
		Code:    127988,
	},
	"envelope": Emoji{
		IconURL: fmt.Sprintf("%s/2709.png?v8", BaseURL),
		Code:    9993,
	},
	"envelope_with_arrow": Emoji{
		IconURL: fmt.Sprintf("%s/1f4e9.png?v8", BaseURL),
		Code:    128233,
	},
	"equatorial_guinea": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1f6.png?v8", BaseURL),
		Code:    127468,
	},
	"eritrea": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ea-1f1f7.png?v8", BaseURL),
		Code:    127466,
	},
	"es": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ea-1f1f8.png?v8", BaseURL),
		Code:    127466,
	},
	"estonia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ea-1f1ea.png?v8", BaseURL),
		Code:    127466,
	},
	"ethiopia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ea-1f1f9.png?v8", BaseURL),
		Code:    127466,
	},
	"eu": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ea-1f1fa.png?v8", BaseURL),
		Code:    127466,
	},
	"euro": Emoji{
		IconURL: fmt.Sprintf("%s/1f4b6.png?v8", BaseURL),
		Code:    128182,
	},
	"european_castle": Emoji{
		IconURL: fmt.Sprintf("%s/1f3f0.png?v8", BaseURL),
		Code:    127984,
	},
	"european_post_office": Emoji{
		IconURL: fmt.Sprintf("%s/1f3e4.png?v8", BaseURL),
		Code:    127972,
	},
	"european_union": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ea-1f1fa.png?v8", BaseURL),
		Code:    127466,
	},
	"evergreen_tree": Emoji{
		IconURL: fmt.Sprintf("%s/1f332.png?v8", BaseURL),
		Code:    127794,
	},
	"exclamation": Emoji{
		IconURL: fmt.Sprintf("%s/2757.png?v8", BaseURL),
		Code:    10071,
	},
	"exploding_head": Emoji{
		IconURL: fmt.Sprintf("%s/1f92f.png?v8", BaseURL),
		Code:    129327,
	},
	"expressionless": Emoji{
		IconURL: fmt.Sprintf("%s/1f611.png?v8", BaseURL),
		Code:    128529,
	},
	"eye": Emoji{
		IconURL: fmt.Sprintf("%s/1f441.png?v8", BaseURL),
		Code:    128065,
	},
	"eye_speech_bubble": Emoji{
		IconURL: fmt.Sprintf("%s/1f441-1f5e8.png?v8", BaseURL),
		Code:    128065,
	},
	"eyeglasses": Emoji{
		IconURL: fmt.Sprintf("%s/1f453.png?v8", BaseURL),
		Code:    128083,
	},
	"eyes": Emoji{
		IconURL: fmt.Sprintf("%s/1f440.png?v8", BaseURL),
		Code:    128064,
	},
	"face_exhaling": Emoji{
		IconURL: fmt.Sprintf("%s/1f62e-1f4a8.png?v8", BaseURL),
		Code:    128558,
	},
	"face_in_clouds": Emoji{
		IconURL: fmt.Sprintf("%s/1f636-1f32b.png?v8", BaseURL),
		Code:    128566,
	},
	"face_with_head_bandage": Emoji{
		IconURL: fmt.Sprintf("%s/1f915.png?v8", BaseURL),
		Code:    129301,
	},
	"face_with_spiral_eyes": Emoji{
		IconURL: fmt.Sprintf("%s/1f635-1f4ab.png?v8", BaseURL),
		Code:    128565,
	},
	"face_with_thermometer": Emoji{
		IconURL: fmt.Sprintf("%s/1f912.png?v8", BaseURL),
		Code:    129298,
	},
	"facepalm": Emoji{
		IconURL: fmt.Sprintf("%s/1f926.png?v8", BaseURL),
		Code:    129318,
	},
	"facepunch": Emoji{
		IconURL: fmt.Sprintf("%s/1f44a.png?v8", BaseURL),
		Code:    128074,
	},
	"factory": Emoji{
		IconURL: fmt.Sprintf("%s/1f3ed.png?v8", BaseURL),
		Code:    127981,
	},
	"factory_worker": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f3ed.png?v8", BaseURL),
		Code:    129489,
	},
	"fairy": Emoji{
		IconURL: fmt.Sprintf("%s/1f9da.png?v8", BaseURL),
		Code:    129498,
	},
	"fairy_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f9da-2642.png?v8", BaseURL),
		Code:    129498,
	},
	"fairy_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f9da-2640.png?v8", BaseURL),
		Code:    129498,
	},
	"falafel": Emoji{
		IconURL: fmt.Sprintf("%s/1f9c6.png?v8", BaseURL),
		Code:    129478,
	},
	"falkland_islands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1eb-1f1f0.png?v8", BaseURL),
		Code:    127467,
	},
	"fallen_leaf": Emoji{
		IconURL: fmt.Sprintf("%s/1f342.png?v8", BaseURL),
		Code:    127810,
	},
	"family": Emoji{
		IconURL: fmt.Sprintf("%s/1f46a.png?v8", BaseURL),
		Code:    128106,
	},
	"family_man_boy": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f466.png?v8", BaseURL),
		Code:    128104,
	},
	"family_man_boy_boy": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f466-1f466.png?v8", BaseURL),
		Code:    128104,
	},
	"family_man_girl": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f467.png?v8", BaseURL),
		Code:    128104,
	},
	"family_man_girl_boy": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f467-1f466.png?v8", BaseURL),
		Code:    128104,
	},
	"family_man_girl_girl": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f467-1f467.png?v8", BaseURL),
		Code:    128104,
	},
	"family_man_man_boy": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f468-1f466.png?v8", BaseURL),
		Code:    128104,
	},
	"family_man_man_boy_boy": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f468-1f466-1f466.png?v8", BaseURL),
		Code:    128104,
	},
	"family_man_man_girl": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f468-1f467.png?v8", BaseURL),
		Code:    128104,
	},
	"family_man_man_girl_boy": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f468-1f467-1f466.png?v8", BaseURL),
		Code:    128104,
	},
	"family_man_man_girl_girl": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f468-1f467-1f467.png?v8", BaseURL),
		Code:    128104,
	},
	"family_man_woman_boy": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f469-1f466.png?v8", BaseURL),
		Code:    128104,
	},
	"family_man_woman_boy_boy": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f469-1f466-1f466.png?v8", BaseURL),
		Code:    128104,
	},
	"family_man_woman_girl": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f469-1f467.png?v8", BaseURL),
		Code:    128104,
	},
	"family_man_woman_girl_boy": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f469-1f467-1f466.png?v8", BaseURL),
		Code:    128104,
	},
	"family_man_woman_girl_girl": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f469-1f467-1f467.png?v8", BaseURL),
		Code:    128104,
	},
	"family_woman_boy": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f466.png?v8", BaseURL),
		Code:    128105,
	},
	"family_woman_boy_boy": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f466-1f466.png?v8", BaseURL),
		Code:    128105,
	},
	"family_woman_girl": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f467.png?v8", BaseURL),
		Code:    128105,
	},
	"family_woman_girl_boy": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f467-1f466.png?v8", BaseURL),
		Code:    128105,
	},
	"family_woman_girl_girl": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f467-1f467.png?v8", BaseURL),
		Code:    128105,
	},
	"family_woman_woman_boy": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f469-1f466.png?v8", BaseURL),
		Code:    128105,
	},
	"family_woman_woman_boy_boy": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f469-1f466-1f466.png?v8", BaseURL),
		Code:    128105,
	},
	"family_woman_woman_girl": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f469-1f467.png?v8", BaseURL),
		Code:    128105,
	},
	"family_woman_woman_girl_boy": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f469-1f467-1f466.png?v8", BaseURL),
		Code:    128105,
	},
	"family_woman_woman_girl_girl": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f469-1f467-1f467.png?v8", BaseURL),
		Code:    128105,
	},
	"farmer": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f33e.png?v8", BaseURL),
		Code:    129489,
	},
	"faroe_islands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1eb-1f1f4.png?v8", BaseURL),
		Code:    127467,
	},
	"fast_forward": Emoji{
		IconURL: fmt.Sprintf("%s/23e9.png?v8", BaseURL),
		Code:    9193,
	},
	"fax": Emoji{
		IconURL: fmt.Sprintf("%s/1f4e0.png?v8", BaseURL),
		Code:    128224,
	},
	"fearful": Emoji{
		IconURL: fmt.Sprintf("%s/1f628.png?v8", BaseURL),
		Code:    128552,
	},
	"feather": Emoji{
		IconURL: fmt.Sprintf("%s/1fab6.png?v8", BaseURL),
		Code:    129718,
	},
	"feelsgood": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/feelsgood.png?v8",
		Code:    0,
	},
	"feet": Emoji{
		IconURL: fmt.Sprintf("%s/1f43e.png?v8", BaseURL),
		Code:    128062,
	},
	"female_detective": Emoji{
		IconURL: fmt.Sprintf("%s/1f575-2640.png?v8", BaseURL),
		Code:    128373,
	},
	"female_sign": Emoji{
		IconURL: fmt.Sprintf("%s/2640.png?v8", BaseURL),
		Code:    9792,
	},
	"ferris_wheel": Emoji{
		IconURL: fmt.Sprintf("%s/1f3a1.png?v8", BaseURL),
		Code:    127905,
	},
	"ferry": Emoji{
		IconURL: fmt.Sprintf("%s/26f4.png?v8", BaseURL),
		Code:    9972,
	},
	"field_hockey": Emoji{
		IconURL: fmt.Sprintf("%s/1f3d1.png?v8", BaseURL),
		Code:    127953,
	},
	"fiji": Emoji{
		IconURL: fmt.Sprintf("%s/1f1eb-1f1ef.png?v8", BaseURL),
		Code:    127467,
	},
	"file_cabinet": Emoji{
		IconURL: fmt.Sprintf("%s/1f5c4.png?v8", BaseURL),
		Code:    128452,
	},
	"file_folder": Emoji{
		IconURL: fmt.Sprintf("%s/1f4c1.png?v8", BaseURL),
		Code:    128193,
	},
	"film_projector": Emoji{
		IconURL: fmt.Sprintf("%s/1f4fd.png?v8", BaseURL),
		Code:    128253,
	},
	"film_strip": Emoji{
		IconURL: fmt.Sprintf("%s/1f39e.png?v8", BaseURL),
		Code:    127902,
	},
	"finland": Emoji{
		IconURL: fmt.Sprintf("%s/1f1eb-1f1ee.png?v8", BaseURL),
		Code:    127467,
	},
	"finnadie": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/finnadie.png?v8",
		Code:    0,
	},
	"fire": Emoji{
		IconURL: fmt.Sprintf("%s/1f525.png?v8", BaseURL),
		Code:    128293,
	},
	"fire_engine": Emoji{
		IconURL: fmt.Sprintf("%s/1f692.png?v8", BaseURL),
		Code:    128658,
	},
	"fire_extinguisher": Emoji{
		IconURL: fmt.Sprintf("%s/1f9ef.png?v8", BaseURL),
		Code:    129519,
	},
	"firecracker": Emoji{
		IconURL: fmt.Sprintf("%s/1f9e8.png?v8", BaseURL),
		Code:    129512,
	},
	"firefighter": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f692.png?v8", BaseURL),
		Code:    129489,
	},
	"fireworks": Emoji{
		IconURL: fmt.Sprintf("%s/1f386.png?v8", BaseURL),
		Code:    127878,
	},
	"first_quarter_moon": Emoji{
		IconURL: fmt.Sprintf("%s/1f313.png?v8", BaseURL),
		Code:    127763,
	},
	"first_quarter_moon_with_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f31b.png?v8", BaseURL),
		Code:    127771,
	},
	"fish": Emoji{
		IconURL: fmt.Sprintf("%s/1f41f.png?v8", BaseURL),
		Code:    128031,
	},
	"fish_cake": Emoji{
		IconURL: fmt.Sprintf("%s/1f365.png?v8", BaseURL),
		Code:    127845,
	},
	"fishing_pole_and_fish": Emoji{
		IconURL: fmt.Sprintf("%s/1f3a3.png?v8", BaseURL),
		Code:    127907,
	},
	"fist": Emoji{
		IconURL: fmt.Sprintf("%s/270a.png?v8", BaseURL),
		Code:    9994,
	},
	"fist_left": Emoji{
		IconURL: fmt.Sprintf("%s/1f91b.png?v8", BaseURL),
		Code:    129307,
	},
	"fist_oncoming": Emoji{
		IconURL: fmt.Sprintf("%s/1f44a.png?v8", BaseURL),
		Code:    128074,
	},
	"fist_raised": Emoji{
		IconURL: fmt.Sprintf("%s/270a.png?v8", BaseURL),
		Code:    9994,
	},
	"fist_right": Emoji{
		IconURL: fmt.Sprintf("%s/1f91c.png?v8", BaseURL),
		Code:    129308,
	},
	"five": Emoji{
		IconURL: fmt.Sprintf("%s/0035-20e3.png?v8", BaseURL),
		Code:    53,
	},
	"flags": Emoji{
		IconURL: fmt.Sprintf("%s/1f38f.png?v8", BaseURL),
		Code:    127887,
	},
	"flamingo": Emoji{
		IconURL: fmt.Sprintf("%s/1f9a9.png?v8", BaseURL),
		Code:    129449,
	},
	"flashlight": Emoji{
		IconURL: fmt.Sprintf("%s/1f526.png?v8", BaseURL),
		Code:    128294,
	},
	"flat_shoe": Emoji{
		IconURL: fmt.Sprintf("%s/1f97f.png?v8", BaseURL),
		Code:    129407,
	},
	"flatbread": Emoji{
		IconURL: fmt.Sprintf("%s/1fad3.png?v8", BaseURL),
		Code:    129747,
	},
	"fleur_de_lis": Emoji{
		IconURL: fmt.Sprintf("%s/269c.png?v8", BaseURL),
		Code:    9884,
	},
	"flight_arrival": Emoji{
		IconURL: fmt.Sprintf("%s/1f6ec.png?v8", BaseURL),
		Code:    128748,
	},
	"flight_departure": Emoji{
		IconURL: fmt.Sprintf("%s/1f6eb.png?v8", BaseURL),
		Code:    128747,
	},
	"flipper": Emoji{
		IconURL: fmt.Sprintf("%s/1f42c.png?v8", BaseURL),
		Code:    128044,
	},
	"floppy_disk": Emoji{
		IconURL: fmt.Sprintf("%s/1f4be.png?v8", BaseURL),
		Code:    128190,
	},
	"flower_playing_cards": Emoji{
		IconURL: fmt.Sprintf("%s/1f3b4.png?v8", BaseURL),
		Code:    127924,
	},
	"flushed": Emoji{
		IconURL: fmt.Sprintf("%s/1f633.png?v8", BaseURL),
		Code:    128563,
	},
	"fly": Emoji{
		IconURL: fmt.Sprintf("%s/1fab0.png?v8", BaseURL),
		Code:    129712,
	},
	"flying_disc": Emoji{
		IconURL: fmt.Sprintf("%s/1f94f.png?v8", BaseURL),
		Code:    129359,
	},
	"flying_saucer": Emoji{
		IconURL: fmt.Sprintf("%s/1f6f8.png?v8", BaseURL),
		Code:    128760,
	},
	"fog": Emoji{
		IconURL: fmt.Sprintf("%s/1f32b.png?v8", BaseURL),
		Code:    127787,
	},
	"foggy": Emoji{
		IconURL: fmt.Sprintf("%s/1f301.png?v8", BaseURL),
		Code:    127745,
	},
	"fondue": Emoji{
		IconURL: fmt.Sprintf("%s/1fad5.png?v8", BaseURL),
		Code:    129749,
	},
	"foot": Emoji{
		IconURL: fmt.Sprintf("%s/1f9b6.png?v8", BaseURL),
		Code:    129462,
	},
	"football": Emoji{
		IconURL: fmt.Sprintf("%s/1f3c8.png?v8", BaseURL),
		Code:    127944,
	},
	"footprints": Emoji{
		IconURL: fmt.Sprintf("%s/1f463.png?v8", BaseURL),
		Code:    128099,
	},
	"fork_and_knife": Emoji{
		IconURL: fmt.Sprintf("%s/1f374.png?v8", BaseURL),
		Code:    127860,
	},
	"fortune_cookie": Emoji{
		IconURL: fmt.Sprintf("%s/1f960.png?v8", BaseURL),
		Code:    129376,
	},
	"fountain": Emoji{
		IconURL: fmt.Sprintf("%s/26f2.png?v8", BaseURL),
		Code:    9970,
	},
	"fountain_pen": Emoji{
		IconURL: fmt.Sprintf("%s/1f58b.png?v8", BaseURL),
		Code:    128395,
	},
	"four": Emoji{
		IconURL: fmt.Sprintf("%s/0034-20e3.png?v8", BaseURL),
		Code:    52,
	},
	"four_leaf_clover": Emoji{
		IconURL: fmt.Sprintf("%s/1f340.png?v8", BaseURL),
		Code:    127808,
	},
	"fox_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f98a.png?v8", BaseURL),
		Code:    129418,
	},
	"fr": Emoji{
		IconURL: fmt.Sprintf("%s/1f1eb-1f1f7.png?v8", BaseURL),
		Code:    127467,
	},
	"framed_picture": Emoji{
		IconURL: fmt.Sprintf("%s/1f5bc.png?v8", BaseURL),
		Code:    128444,
	},
	"free": Emoji{
		IconURL: fmt.Sprintf("%s/1f193.png?v8", BaseURL),
		Code:    127379,
	},
	"french_guiana": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1eb.png?v8", BaseURL),
		Code:    127468,
	},
	"french_polynesia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f5-1f1eb.png?v8", BaseURL),
		Code:    127477,
	},
	"french_southern_territories": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f9-1f1eb.png?v8", BaseURL),
		Code:    127481,
	},
	"fried_egg": Emoji{
		IconURL: fmt.Sprintf("%s/1f373.png?v8", BaseURL),
		Code:    127859,
	},
	"fried_shrimp": Emoji{
		IconURL: fmt.Sprintf("%s/1f364.png?v8", BaseURL),
		Code:    127844,
	},
	"fries": Emoji{
		IconURL: fmt.Sprintf("%s/1f35f.png?v8", BaseURL),
		Code:    127839,
	},
	"frog": Emoji{
		IconURL: fmt.Sprintf("%s/1f438.png?v8", BaseURL),
		Code:    128056,
	},
	"frowning": Emoji{
		IconURL: fmt.Sprintf("%s/1f626.png?v8", BaseURL),
		Code:    128550,
	},
	"frowning_face": Emoji{
		IconURL: fmt.Sprintf("%s/2639.png?v8", BaseURL),
		Code:    9785,
	},
	"frowning_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f64d-2642.png?v8", BaseURL),
		Code:    128589,
	},
	"frowning_person": Emoji{
		IconURL: fmt.Sprintf("%s/1f64d.png?v8", BaseURL),
		Code:    128589,
	},
	"frowning_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f64d-2640.png?v8", BaseURL),
		Code:    128589,
	},
	"fu": Emoji{
		IconURL: fmt.Sprintf("%s/1f595.png?v8", BaseURL),
		Code:    128405,
	},
	"fuelpump": Emoji{
		IconURL: fmt.Sprintf("%s/26fd.png?v8", BaseURL),
		Code:    9981,
	},
	"full_moon": Emoji{
		IconURL: fmt.Sprintf("%s/1f315.png?v8", BaseURL),
		Code:    127765,
	},
	"full_moon_with_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f31d.png?v8", BaseURL),
		Code:    127773,
	},
	"funeral_urn": Emoji{
		IconURL: fmt.Sprintf("%s/26b1.png?v8", BaseURL),
		Code:    9905,
	},
	"gabon": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1e6.png?v8", BaseURL),
		Code:    127468,
	},
	"gambia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1f2.png?v8", BaseURL),
		Code:    127468,
	},
	"game_die": Emoji{
		IconURL: fmt.Sprintf("%s/1f3b2.png?v8", BaseURL),
		Code:    127922,
	},
	"garlic": Emoji{
		IconURL: fmt.Sprintf("%s/1f9c4.png?v8", BaseURL),
		Code:    129476,
	},
	"gb": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1e7.png?v8", BaseURL),
		Code:    127468,
	},
	"gear": Emoji{
		IconURL: fmt.Sprintf("%s/2699.png?v8", BaseURL),
		Code:    9881,
	},
	"gem": Emoji{
		IconURL: fmt.Sprintf("%s/1f48e.png?v8", BaseURL),
		Code:    128142,
	},
	"gemini": Emoji{
		IconURL: fmt.Sprintf("%s/264a.png?v8", BaseURL),
		Code:    9802,
	},
	"genie": Emoji{
		IconURL: fmt.Sprintf("%s/1f9de.png?v8", BaseURL),
		Code:    129502,
	},
	"genie_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f9de-2642.png?v8", BaseURL),
		Code:    129502,
	},
	"genie_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f9de-2640.png?v8", BaseURL),
		Code:    129502,
	},
	"georgia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1ea.png?v8", BaseURL),
		Code:    127468,
	},
	"ghana": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1ed.png?v8", BaseURL),
		Code:    127468,
	},
	"ghost": Emoji{
		IconURL: fmt.Sprintf("%s/1f47b.png?v8", BaseURL),
		Code:    128123,
	},
	"gibraltar": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1ee.png?v8", BaseURL),
		Code:    127468,
	},
	"gift": Emoji{
		IconURL: fmt.Sprintf("%s/1f381.png?v8", BaseURL),
		Code:    127873,
	},
	"gift_heart": Emoji{
		IconURL: fmt.Sprintf("%s/1f49d.png?v8", BaseURL),
		Code:    128157,
	},
	"giraffe": Emoji{
		IconURL: fmt.Sprintf("%s/1f992.png?v8", BaseURL),
		Code:    129426,
	},
	"girl": Emoji{
		IconURL: fmt.Sprintf("%s/1f467.png?v8", BaseURL),
		Code:    128103,
	},
	"globe_with_meridians": Emoji{
		IconURL: fmt.Sprintf("%s/1f310.png?v8", BaseURL),
		Code:    127760,
	},
	"gloves": Emoji{
		IconURL: fmt.Sprintf("%s/1f9e4.png?v8", BaseURL),
		Code:    129508,
	},
	"goal_net": Emoji{
		IconURL: fmt.Sprintf("%s/1f945.png?v8", BaseURL),
		Code:    129349,
	},
	"goat": Emoji{
		IconURL: fmt.Sprintf("%s/1f410.png?v8", BaseURL),
		Code:    128016,
	},
	"goberserk": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/goberserk.png?v8",
		Code:    0,
	},
	"godmode": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/godmode.png?v8",
		Code:    0,
	},
	"goggles": Emoji{
		IconURL: fmt.Sprintf("%s/1f97d.png?v8", BaseURL),
		Code:    129405,
	},
	"golf": Emoji{
		IconURL: fmt.Sprintf("%s/26f3.png?v8", BaseURL),
		Code:    9971,
	},
	"golfing": Emoji{
		IconURL: fmt.Sprintf("%s/1f3cc.png?v8", BaseURL),
		Code:    127948,
	},
	"golfing_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f3cc-2642.png?v8", BaseURL),
		Code:    127948,
	},
	"golfing_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f3cc-2640.png?v8", BaseURL),
		Code:    127948,
	},
	"gorilla": Emoji{
		IconURL: fmt.Sprintf("%s/1f98d.png?v8", BaseURL),
		Code:    129421,
	},
	"grapes": Emoji{
		IconURL: fmt.Sprintf("%s/1f347.png?v8", BaseURL),
		Code:    127815,
	},
	"greece": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1f7.png?v8", BaseURL),
		Code:    127468,
	},
	"green_apple": Emoji{
		IconURL: fmt.Sprintf("%s/1f34f.png?v8", BaseURL),
		Code:    127823,
	},
	"green_book": Emoji{
		IconURL: fmt.Sprintf("%s/1f4d7.png?v8", BaseURL),
		Code:    128215,
	},
	"green_circle": Emoji{
		IconURL: fmt.Sprintf("%s/1f7e2.png?v8", BaseURL),
		Code:    128994,
	},
	"green_heart": Emoji{
		IconURL: fmt.Sprintf("%s/1f49a.png?v8", BaseURL),
		Code:    128154,
	},
	"green_salad": Emoji{
		IconURL: fmt.Sprintf("%s/1f957.png?v8", BaseURL),
		Code:    129367,
	},
	"green_square": Emoji{
		IconURL: fmt.Sprintf("%s/1f7e9.png?v8", BaseURL),
		Code:    129001,
	},
	"greenland": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1f1.png?v8", BaseURL),
		Code:    127468,
	},
	"grenada": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1e9.png?v8", BaseURL),
		Code:    127468,
	},
	"grey_exclamation": Emoji{
		IconURL: fmt.Sprintf("%s/2755.png?v8", BaseURL),
		Code:    10069,
	},
	"grey_question": Emoji{
		IconURL: fmt.Sprintf("%s/2754.png?v8", BaseURL),
		Code:    10068,
	},
	"grimacing": Emoji{
		IconURL: fmt.Sprintf("%s/1f62c.png?v8", BaseURL),
		Code:    128556,
	},
	"grin": Emoji{
		IconURL: fmt.Sprintf("%s/1f601.png?v8", BaseURL),
		Code:    128513,
	},
	"grinning": Emoji{
		IconURL: fmt.Sprintf("%s/1f600.png?v8", BaseURL),
		Code:    128512,
	},
	"guadeloupe": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1f5.png?v8", BaseURL),
		Code:    127468,
	},
	"guam": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1fa.png?v8", BaseURL),
		Code:    127468,
	},
	"guard": Emoji{
		IconURL: fmt.Sprintf("%s/1f482.png?v8", BaseURL),
		Code:    128130,
	},
	"guardsman": Emoji{
		IconURL: fmt.Sprintf("%s/1f482-2642.png?v8", BaseURL),
		Code:    128130,
	},
	"guardswoman": Emoji{
		IconURL: fmt.Sprintf("%s/1f482-2640.png?v8", BaseURL),
		Code:    128130,
	},
	"guatemala": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1f9.png?v8", BaseURL),
		Code:    127468,
	},
	"guernsey": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1ec.png?v8", BaseURL),
		Code:    127468,
	},
	"guide_dog": Emoji{
		IconURL: fmt.Sprintf("%s/1f9ae.png?v8", BaseURL),
		Code:    129454,
	},
	"guinea": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1f3.png?v8", BaseURL),
		Code:    127468,
	},
	"guinea_bissau": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1fc.png?v8", BaseURL),
		Code:    127468,
	},
	"guitar": Emoji{
		IconURL: fmt.Sprintf("%s/1f3b8.png?v8", BaseURL),
		Code:    127928,
	},
	"gun": Emoji{
		IconURL: fmt.Sprintf("%s/1f52b.png?v8", BaseURL),
		Code:    128299,
	},
	"guyana": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1fe.png?v8", BaseURL),
		Code:    127468,
	},
	"haircut": Emoji{
		IconURL: fmt.Sprintf("%s/1f487.png?v8", BaseURL),
		Code:    128135,
	},
	"haircut_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f487-2642.png?v8", BaseURL),
		Code:    128135,
	},
	"haircut_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f487-2640.png?v8", BaseURL),
		Code:    128135,
	},
	"haiti": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ed-1f1f9.png?v8", BaseURL),
		Code:    127469,
	},
	"hamburger": Emoji{
		IconURL: fmt.Sprintf("%s/1f354.png?v8", BaseURL),
		Code:    127828,
	},
	"hammer": Emoji{
		IconURL: fmt.Sprintf("%s/1f528.png?v8", BaseURL),
		Code:    128296,
	},
	"hammer_and_pick": Emoji{
		IconURL: fmt.Sprintf("%s/2692.png?v8", BaseURL),
		Code:    9874,
	},
	"hammer_and_wrench": Emoji{
		IconURL: fmt.Sprintf("%s/1f6e0.png?v8", BaseURL),
		Code:    128736,
	},
	"hamster": Emoji{
		IconURL: fmt.Sprintf("%s/1f439.png?v8", BaseURL),
		Code:    128057,
	},
	"hand": Emoji{
		IconURL: fmt.Sprintf("%s/270b.png?v8", BaseURL),
		Code:    9995,
	},
	"hand_over_mouth": Emoji{
		IconURL: fmt.Sprintf("%s/1f92d.png?v8", BaseURL),
		Code:    129325,
	},
	"handbag": Emoji{
		IconURL: fmt.Sprintf("%s/1f45c.png?v8", BaseURL),
		Code:    128092,
	},
	"handball_person": Emoji{
		IconURL: fmt.Sprintf("%s/1f93e.png?v8", BaseURL),
		Code:    129342,
	},
	"handshake": Emoji{
		IconURL: fmt.Sprintf("%s/1f91d.png?v8", BaseURL),
		Code:    129309,
	},
	"hankey": Emoji{
		IconURL: fmt.Sprintf("%s/1f4a9.png?v8", BaseURL),
		Code:    128169,
	},
	"hash": Emoji{
		IconURL: fmt.Sprintf("%s/0023-20e3.png?v8", BaseURL),
		Code:    35,
	},
	"hatched_chick": Emoji{
		IconURL: fmt.Sprintf("%s/1f425.png?v8", BaseURL),
		Code:    128037,
	},
	"hatching_chick": Emoji{
		IconURL: fmt.Sprintf("%s/1f423.png?v8", BaseURL),
		Code:    128035,
	},
	"headphones": Emoji{
		IconURL: fmt.Sprintf("%s/1f3a7.png?v8", BaseURL),
		Code:    127911,
	},
	"headstone": Emoji{
		IconURL: fmt.Sprintf("%s/1faa6.png?v8", BaseURL),
		Code:    129702,
	},
	"health_worker": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-2695.png?v8", BaseURL),
		Code:    129489,
	},
	"hear_no_evil": Emoji{
		IconURL: fmt.Sprintf("%s/1f649.png?v8", BaseURL),
		Code:    128585,
	},
	"heard_mcdonald_islands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ed-1f1f2.png?v8", BaseURL),
		Code:    127469,
	},
	"heart": Emoji{
		IconURL: fmt.Sprintf("%s/2764.png?v8", BaseURL),
		Code:    10084,
	},
	"heart_decoration": Emoji{
		IconURL: fmt.Sprintf("%s/1f49f.png?v8", BaseURL),
		Code:    128159,
	},
	"heart_eyes": Emoji{
		IconURL: fmt.Sprintf("%s/1f60d.png?v8", BaseURL),
		Code:    128525,
	},
	"heart_eyes_cat": Emoji{
		IconURL: fmt.Sprintf("%s/1f63b.png?v8", BaseURL),
		Code:    128571,
	},
	"heart_on_fire": Emoji{
		IconURL: fmt.Sprintf("%s/2764-1f525.png?v8", BaseURL),
		Code:    10084,
	},
	"heartbeat": Emoji{
		IconURL: fmt.Sprintf("%s/1f493.png?v8", BaseURL),
		Code:    128147,
	},
	"heartpulse": Emoji{
		IconURL: fmt.Sprintf("%s/1f497.png?v8", BaseURL),
		Code:    128151,
	},
	"hearts": Emoji{
		IconURL: fmt.Sprintf("%s/2665.png?v8", BaseURL),
		Code:    9829,
	},
	"heavy_check_mark": Emoji{
		IconURL: fmt.Sprintf("%s/2714.png?v8", BaseURL),
		Code:    10004,
	},
	"heavy_division_sign": Emoji{
		IconURL: fmt.Sprintf("%s/2797.png?v8", BaseURL),
		Code:    10135,
	},
	"heavy_dollar_sign": Emoji{
		IconURL: fmt.Sprintf("%s/1f4b2.png?v8", BaseURL),
		Code:    128178,
	},
	"heavy_exclamation_mark": Emoji{
		IconURL: fmt.Sprintf("%s/2757.png?v8", BaseURL),
		Code:    10071,
	},
	"heavy_heart_exclamation": Emoji{
		IconURL: fmt.Sprintf("%s/2763.png?v8", BaseURL),
		Code:    10083,
	},
	"heavy_minus_sign": Emoji{
		IconURL: fmt.Sprintf("%s/2796.png?v8", BaseURL),
		Code:    10134,
	},
	"heavy_multiplication_x": Emoji{
		IconURL: fmt.Sprintf("%s/2716.png?v8", BaseURL),
		Code:    10006,
	},
	"heavy_plus_sign": Emoji{
		IconURL: fmt.Sprintf("%s/2795.png?v8", BaseURL),
		Code:    10133,
	},
	"hedgehog": Emoji{
		IconURL: fmt.Sprintf("%s/1f994.png?v8", BaseURL),
		Code:    129428,
	},
	"helicopter": Emoji{
		IconURL: fmt.Sprintf("%s/1f681.png?v8", BaseURL),
		Code:    128641,
	},
	"herb": Emoji{
		IconURL: fmt.Sprintf("%s/1f33f.png?v8", BaseURL),
		Code:    127807,
	},
	"hibiscus": Emoji{
		IconURL: fmt.Sprintf("%s/1f33a.png?v8", BaseURL),
		Code:    127802,
	},
	"high_brightness": Emoji{
		IconURL: fmt.Sprintf("%s/1f506.png?v8", BaseURL),
		Code:    128262,
	},
	"high_heel": Emoji{
		IconURL: fmt.Sprintf("%s/1f460.png?v8", BaseURL),
		Code:    128096,
	},
	"hiking_boot": Emoji{
		IconURL: fmt.Sprintf("%s/1f97e.png?v8", BaseURL),
		Code:    129406,
	},
	"hindu_temple": Emoji{
		IconURL: fmt.Sprintf("%s/1f6d5.png?v8", BaseURL),
		Code:    128725,
	},
	"hippopotamus": Emoji{
		IconURL: fmt.Sprintf("%s/1f99b.png?v8", BaseURL),
		Code:    129435,
	},
	"hocho": Emoji{
		IconURL: fmt.Sprintf("%s/1f52a.png?v8", BaseURL),
		Code:    128298,
	},
	"hole": Emoji{
		IconURL: fmt.Sprintf("%s/1f573.png?v8", BaseURL),
		Code:    128371,
	},
	"honduras": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ed-1f1f3.png?v8", BaseURL),
		Code:    127469,
	},
	"honey_pot": Emoji{
		IconURL: fmt.Sprintf("%s/1f36f.png?v8", BaseURL),
		Code:    127855,
	},
	"honeybee": Emoji{
		IconURL: fmt.Sprintf("%s/1f41d.png?v8", BaseURL),
		Code:    128029,
	},
	"hong_kong": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ed-1f1f0.png?v8", BaseURL),
		Code:    127469,
	},
	"hook": Emoji{
		IconURL: fmt.Sprintf("%s/1fa9d.png?v8", BaseURL),
		Code:    129693,
	},
	"horse": Emoji{
		IconURL: fmt.Sprintf("%s/1f434.png?v8", BaseURL),
		Code:    128052,
	},
	"horse_racing": Emoji{
		IconURL: fmt.Sprintf("%s/1f3c7.png?v8", BaseURL),
		Code:    127943,
	},
	"hospital": Emoji{
		IconURL: fmt.Sprintf("%s/1f3e5.png?v8", BaseURL),
		Code:    127973,
	},
	"hot_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f975.png?v8", BaseURL),
		Code:    129397,
	},
	"hot_pepper": Emoji{
		IconURL: fmt.Sprintf("%s/1f336.png?v8", BaseURL),
		Code:    127798,
	},
	"hotdog": Emoji{
		IconURL: fmt.Sprintf("%s/1f32d.png?v8", BaseURL),
		Code:    127789,
	},
	"hotel": Emoji{
		IconURL: fmt.Sprintf("%s/1f3e8.png?v8", BaseURL),
		Code:    127976,
	},
	"hotsprings": Emoji{
		IconURL: fmt.Sprintf("%s/2668.png?v8", BaseURL),
		Code:    9832,
	},
	"hourglass": Emoji{
		IconURL: fmt.Sprintf("%s/231b.png?v8", BaseURL),
		Code:    8987,
	},
	"hourglass_flowing_sand": Emoji{
		IconURL: fmt.Sprintf("%s/23f3.png?v8", BaseURL),
		Code:    9203,
	},
	"house": Emoji{
		IconURL: fmt.Sprintf("%s/1f3e0.png?v8", BaseURL),
		Code:    127968,
	},
	"house_with_garden": Emoji{
		IconURL: fmt.Sprintf("%s/1f3e1.png?v8", BaseURL),
		Code:    127969,
	},
	"houses": Emoji{
		IconURL: fmt.Sprintf("%s/1f3d8.png?v8", BaseURL),
		Code:    127960,
	},
	"hugs": Emoji{
		IconURL: fmt.Sprintf("%s/1f917.png?v8", BaseURL),
		Code:    129303,
	},
	"hungary": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ed-1f1fa.png?v8", BaseURL),
		Code:    127469,
	},
	"hurtrealbad": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/hurtrealbad.png?v8",
		Code:    0,
	},
	"hushed": Emoji{
		IconURL: fmt.Sprintf("%s/1f62f.png?v8", BaseURL),
		Code:    128559,
	},
	"hut": Emoji{
		IconURL: fmt.Sprintf("%s/1f6d6.png?v8", BaseURL),
		Code:    128726,
	},
	"ice_cream": Emoji{
		IconURL: fmt.Sprintf("%s/1f368.png?v8", BaseURL),
		Code:    127848,
	},
	"ice_cube": Emoji{
		IconURL: fmt.Sprintf("%s/1f9ca.png?v8", BaseURL),
		Code:    129482,
	},
	"ice_hockey": Emoji{
		IconURL: fmt.Sprintf("%s/1f3d2.png?v8", BaseURL),
		Code:    127954,
	},
	"ice_skate": Emoji{
		IconURL: fmt.Sprintf("%s/26f8.png?v8", BaseURL),
		Code:    9976,
	},
	"icecream": Emoji{
		IconURL: fmt.Sprintf("%s/1f366.png?v8", BaseURL),
		Code:    127846,
	},
	"iceland": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ee-1f1f8.png?v8", BaseURL),
		Code:    127470,
	},
	"id": Emoji{
		IconURL: fmt.Sprintf("%s/1f194.png?v8", BaseURL),
		Code:    127380,
	},
	"ideograph_advantage": Emoji{
		IconURL: fmt.Sprintf("%s/1f250.png?v8", BaseURL),
		Code:    127568,
	},
	"imp": Emoji{
		IconURL: fmt.Sprintf("%s/1f47f.png?v8", BaseURL),
		Code:    128127,
	},
	"inbox_tray": Emoji{
		IconURL: fmt.Sprintf("%s/1f4e5.png?v8", BaseURL),
		Code:    128229,
	},
	"incoming_envelope": Emoji{
		IconURL: fmt.Sprintf("%s/1f4e8.png?v8", BaseURL),
		Code:    128232,
	},
	"india": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ee-1f1f3.png?v8", BaseURL),
		Code:    127470,
	},
	"indonesia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ee-1f1e9.png?v8", BaseURL),
		Code:    127470,
	},
	"infinity": Emoji{
		IconURL: fmt.Sprintf("%s/267e.png?v8", BaseURL),
		Code:    9854,
	},
	"information_desk_person": Emoji{
		IconURL: fmt.Sprintf("%s/1f481.png?v8", BaseURL),
		Code:    128129,
	},
	"information_source": Emoji{
		IconURL: fmt.Sprintf("%s/2139.png?v8", BaseURL),
		Code:    8505,
	},
	"innocent": Emoji{
		IconURL: fmt.Sprintf("%s/1f607.png?v8", BaseURL),
		Code:    128519,
	},
	"interrobang": Emoji{
		IconURL: fmt.Sprintf("%s/2049.png?v8", BaseURL),
		Code:    8265,
	},
	"iphone": Emoji{
		IconURL: fmt.Sprintf("%s/1f4f1.png?v8", BaseURL),
		Code:    128241,
	},
	"iran": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ee-1f1f7.png?v8", BaseURL),
		Code:    127470,
	},
	"iraq": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ee-1f1f6.png?v8", BaseURL),
		Code:    127470,
	},
	"ireland": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ee-1f1ea.png?v8", BaseURL),
		Code:    127470,
	},
	"isle_of_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ee-1f1f2.png?v8", BaseURL),
		Code:    127470,
	},
	"israel": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ee-1f1f1.png?v8", BaseURL),
		Code:    127470,
	},
	"it": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ee-1f1f9.png?v8", BaseURL),
		Code:    127470,
	},
	"izakaya_lantern": Emoji{
		IconURL: fmt.Sprintf("%s/1f3ee.png?v8", BaseURL),
		Code:    127982,
	},
	"jack_o_lantern": Emoji{
		IconURL: fmt.Sprintf("%s/1f383.png?v8", BaseURL),
		Code:    127875,
	},
	"jamaica": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ef-1f1f2.png?v8", BaseURL),
		Code:    127471,
	},
	"japan": Emoji{
		IconURL: fmt.Sprintf("%s/1f5fe.png?v8", BaseURL),
		Code:    128510,
	},
	"japanese_castle": Emoji{
		IconURL: fmt.Sprintf("%s/1f3ef.png?v8", BaseURL),
		Code:    127983,
	},
	"japanese_goblin": Emoji{
		IconURL: fmt.Sprintf("%s/1f47a.png?v8", BaseURL),
		Code:    128122,
	},
	"japanese_ogre": Emoji{
		IconURL: fmt.Sprintf("%s/1f479.png?v8", BaseURL),
		Code:    128121,
	},
	"jeans": Emoji{
		IconURL: fmt.Sprintf("%s/1f456.png?v8", BaseURL),
		Code:    128086,
	},
	"jersey": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ef-1f1ea.png?v8", BaseURL),
		Code:    127471,
	},
	"jigsaw": Emoji{
		IconURL: fmt.Sprintf("%s/1f9e9.png?v8", BaseURL),
		Code:    129513,
	},
	"jordan": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ef-1f1f4.png?v8", BaseURL),
		Code:    127471,
	},
	"joy": Emoji{
		IconURL: fmt.Sprintf("%s/1f602.png?v8", BaseURL),
		Code:    128514,
	},
	"joy_cat": Emoji{
		IconURL: fmt.Sprintf("%s/1f639.png?v8", BaseURL),
		Code:    128569,
	},
	"joystick": Emoji{
		IconURL: fmt.Sprintf("%s/1f579.png?v8", BaseURL),
		Code:    128377,
	},
	"jp": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ef-1f1f5.png?v8", BaseURL),
		Code:    127471,
	},
	"judge": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-2696.png?v8", BaseURL),
		Code:    129489,
	},
	"juggling_person": Emoji{
		IconURL: fmt.Sprintf("%s/1f939.png?v8", BaseURL),
		Code:    129337,
	},
	"kaaba": Emoji{
		IconURL: fmt.Sprintf("%s/1f54b.png?v8", BaseURL),
		Code:    128331,
	},
	"kangaroo": Emoji{
		IconURL: fmt.Sprintf("%s/1f998.png?v8", BaseURL),
		Code:    129432,
	},
	"kazakhstan": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f0-1f1ff.png?v8", BaseURL),
		Code:    127472,
	},
	"kenya": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f0-1f1ea.png?v8", BaseURL),
		Code:    127472,
	},
	"key": Emoji{
		IconURL: fmt.Sprintf("%s/1f511.png?v8", BaseURL),
		Code:    128273,
	},
	"keyboard": Emoji{
		IconURL: fmt.Sprintf("%s/2328.png?v8", BaseURL),
		Code:    9000,
	},
	"keycap_ten": Emoji{
		IconURL: fmt.Sprintf("%s/1f51f.png?v8", BaseURL),
		Code:    128287,
	},
	"kick_scooter": Emoji{
		IconURL: fmt.Sprintf("%s/1f6f4.png?v8", BaseURL),
		Code:    128756,
	},
	"kimono": Emoji{
		IconURL: fmt.Sprintf("%s/1f458.png?v8", BaseURL),
		Code:    128088,
	},
	"kiribati": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f0-1f1ee.png?v8", BaseURL),
		Code:    127472,
	},
	"kiss": Emoji{
		IconURL: fmt.Sprintf("%s/1f48b.png?v8", BaseURL),
		Code:    128139,
	},
	"kissing": Emoji{
		IconURL: fmt.Sprintf("%s/1f617.png?v8", BaseURL),
		Code:    128535,
	},
	"kissing_cat": Emoji{
		IconURL: fmt.Sprintf("%s/1f63d.png?v8", BaseURL),
		Code:    128573,
	},
	"kissing_closed_eyes": Emoji{
		IconURL: fmt.Sprintf("%s/1f61a.png?v8", BaseURL),
		Code:    128538,
	},
	"kissing_heart": Emoji{
		IconURL: fmt.Sprintf("%s/1f618.png?v8", BaseURL),
		Code:    128536,
	},
	"kissing_smiling_eyes": Emoji{
		IconURL: fmt.Sprintf("%s/1f619.png?v8", BaseURL),
		Code:    128537,
	},
	"kite": Emoji{
		IconURL: fmt.Sprintf("%s/1fa81.png?v8", BaseURL),
		Code:    129665,
	},
	"kiwi_fruit": Emoji{
		IconURL: fmt.Sprintf("%s/1f95d.png?v8", BaseURL),
		Code:    129373,
	},
	"kneeling_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f9ce-2642.png?v8", BaseURL),
		Code:    129486,
	},
	"kneeling_person": Emoji{
		IconURL: fmt.Sprintf("%s/1f9ce.png?v8", BaseURL),
		Code:    129486,
	},
	"kneeling_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f9ce-2640.png?v8", BaseURL),
		Code:    129486,
	},
	"knife": Emoji{
		IconURL: fmt.Sprintf("%s/1f52a.png?v8", BaseURL),
		Code:    128298,
	},
	"knot": Emoji{
		IconURL: fmt.Sprintf("%s/1faa2.png?v8", BaseURL),
		Code:    129698,
	},
	"koala": Emoji{
		IconURL: fmt.Sprintf("%s/1f428.png?v8", BaseURL),
		Code:    128040,
	},
	"koko": Emoji{
		IconURL: fmt.Sprintf("%s/1f201.png?v8", BaseURL),
		Code:    127489,
	},
	"kosovo": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fd-1f1f0.png?v8", BaseURL),
		Code:    127485,
	},
	"kr": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f0-1f1f7.png?v8", BaseURL),
		Code:    127472,
	},
	"kuwait": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f0-1f1fc.png?v8", BaseURL),
		Code:    127472,
	},
	"kyrgyzstan": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f0-1f1ec.png?v8", BaseURL),
		Code:    127472,
	},
	"lab_coat": Emoji{
		IconURL: fmt.Sprintf("%s/1f97c.png?v8", BaseURL),
		Code:    129404,
	},
	"label": Emoji{
		IconURL: fmt.Sprintf("%s/1f3f7.png?v8", BaseURL),
		Code:    127991,
	},
	"lacrosse": Emoji{
		IconURL: fmt.Sprintf("%s/1f94d.png?v8", BaseURL),
		Code:    129357,
	},
	"ladder": Emoji{
		IconURL: fmt.Sprintf("%s/1fa9c.png?v8", BaseURL),
		Code:    129692,
	},
	"lady_beetle": Emoji{
		IconURL: fmt.Sprintf("%s/1f41e.png?v8", BaseURL),
		Code:    128030,
	},
	"lantern": Emoji{
		IconURL: fmt.Sprintf("%s/1f3ee.png?v8", BaseURL),
		Code:    127982,
	},
	"laos": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f1-1f1e6.png?v8", BaseURL),
		Code:    127473,
	},
	"large_blue_circle": Emoji{
		IconURL: fmt.Sprintf("%s/1f535.png?v8", BaseURL),
		Code:    128309,
	},
	"large_blue_diamond": Emoji{
		IconURL: fmt.Sprintf("%s/1f537.png?v8", BaseURL),
		Code:    128311,
	},
	"large_orange_diamond": Emoji{
		IconURL: fmt.Sprintf("%s/1f536.png?v8", BaseURL),
		Code:    128310,
	},
	"last_quarter_moon": Emoji{
		IconURL: fmt.Sprintf("%s/1f317.png?v8", BaseURL),
		Code:    127767,
	},
	"last_quarter_moon_with_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f31c.png?v8", BaseURL),
		Code:    127772,
	},
	"latin_cross": Emoji{
		IconURL: fmt.Sprintf("%s/271d.png?v8", BaseURL),
		Code:    10013,
	},
	"latvia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f1-1f1fb.png?v8", BaseURL),
		Code:    127473,
	},
	"laughing": Emoji{
		IconURL: fmt.Sprintf("%s/1f606.png?v8", BaseURL),
		Code:    128518,
	},
	"leafy_green": Emoji{
		IconURL: fmt.Sprintf("%s/1f96c.png?v8", BaseURL),
		Code:    129388,
	},
	"leaves": Emoji{
		IconURL: fmt.Sprintf("%s/1f343.png?v8", BaseURL),
		Code:    127811,
	},
	"lebanon": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f1-1f1e7.png?v8", BaseURL),
		Code:    127473,
	},
	"ledger": Emoji{
		IconURL: fmt.Sprintf("%s/1f4d2.png?v8", BaseURL),
		Code:    128210,
	},
	"left_luggage": Emoji{
		IconURL: fmt.Sprintf("%s/1f6c5.png?v8", BaseURL),
		Code:    128709,
	},
	"left_right_arrow": Emoji{
		IconURL: fmt.Sprintf("%s/2194.png?v8", BaseURL),
		Code:    8596,
	},
	"left_speech_bubble": Emoji{
		IconURL: fmt.Sprintf("%s/1f5e8.png?v8", BaseURL),
		Code:    128488,
	},
	"leftwards_arrow_with_hook": Emoji{
		IconURL: fmt.Sprintf("%s/21a9.png?v8", BaseURL),
		Code:    8617,
	},
	"leg": Emoji{
		IconURL: fmt.Sprintf("%s/1f9b5.png?v8", BaseURL),
		Code:    129461,
	},
	"lemon": Emoji{
		IconURL: fmt.Sprintf("%s/1f34b.png?v8", BaseURL),
		Code:    127819,
	},
	"leo": Emoji{
		IconURL: fmt.Sprintf("%s/264c.png?v8", BaseURL),
		Code:    9804,
	},
	"leopard": Emoji{
		IconURL: fmt.Sprintf("%s/1f406.png?v8", BaseURL),
		Code:    128006,
	},
	"lesotho": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f1-1f1f8.png?v8", BaseURL),
		Code:    127473,
	},
	"level_slider": Emoji{
		IconURL: fmt.Sprintf("%s/1f39a.png?v8", BaseURL),
		Code:    127898,
	},
	"liberia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f1-1f1f7.png?v8", BaseURL),
		Code:    127473,
	},
	"libra": Emoji{
		IconURL: fmt.Sprintf("%s/264e.png?v8", BaseURL),
		Code:    9806,
	},
	"libya": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f1-1f1fe.png?v8", BaseURL),
		Code:    127473,
	},
	"liechtenstein": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f1-1f1ee.png?v8", BaseURL),
		Code:    127473,
	},
	"light_rail": Emoji{
		IconURL: fmt.Sprintf("%s/1f688.png?v8", BaseURL),
		Code:    128648,
	},
	"link": Emoji{
		IconURL: fmt.Sprintf("%s/1f517.png?v8", BaseURL),
		Code:    128279,
	},
	"lion": Emoji{
		IconURL: fmt.Sprintf("%s/1f981.png?v8", BaseURL),
		Code:    129409,
	},
	"lips": Emoji{
		IconURL: fmt.Sprintf("%s/1f444.png?v8", BaseURL),
		Code:    128068,
	},
	"lipstick": Emoji{
		IconURL: fmt.Sprintf("%s/1f484.png?v8", BaseURL),
		Code:    128132,
	},
	"lithuania": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f1-1f1f9.png?v8", BaseURL),
		Code:    127473,
	},
	"lizard": Emoji{
		IconURL: fmt.Sprintf("%s/1f98e.png?v8", BaseURL),
		Code:    129422,
	},
	"llama": Emoji{
		IconURL: fmt.Sprintf("%s/1f999.png?v8", BaseURL),
		Code:    129433,
	},
	"lobster": Emoji{
		IconURL: fmt.Sprintf("%s/1f99e.png?v8", BaseURL),
		Code:    129438,
	},
	"lock": Emoji{
		IconURL: fmt.Sprintf("%s/1f512.png?v8", BaseURL),
		Code:    128274,
	},
	"lock_with_ink_pen": Emoji{
		IconURL: fmt.Sprintf("%s/1f50f.png?v8", BaseURL),
		Code:    128271,
	},
	"lollipop": Emoji{
		IconURL: fmt.Sprintf("%s/1f36d.png?v8", BaseURL),
		Code:    127853,
	},
	"long_drum": Emoji{
		IconURL: fmt.Sprintf("%s/1fa98.png?v8", BaseURL),
		Code:    129688,
	},
	"loop": Emoji{
		IconURL: fmt.Sprintf("%s/27bf.png?v8", BaseURL),
		Code:    10175,
	},
	"lotion_bottle": Emoji{
		IconURL: fmt.Sprintf("%s/1f9f4.png?v8", BaseURL),
		Code:    129524,
	},
	"lotus_position": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d8.png?v8", BaseURL),
		Code:    129496,
	},
	"lotus_position_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d8-2642.png?v8", BaseURL),
		Code:    129496,
	},
	"lotus_position_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d8-2640.png?v8", BaseURL),
		Code:    129496,
	},
	"loud_sound": Emoji{
		IconURL: fmt.Sprintf("%s/1f50a.png?v8", BaseURL),
		Code:    128266,
	},
	"loudspeaker": Emoji{
		IconURL: fmt.Sprintf("%s/1f4e2.png?v8", BaseURL),
		Code:    128226,
	},
	"love_hotel": Emoji{
		IconURL: fmt.Sprintf("%s/1f3e9.png?v8", BaseURL),
		Code:    127977,
	},
	"love_letter": Emoji{
		IconURL: fmt.Sprintf("%s/1f48c.png?v8", BaseURL),
		Code:    128140,
	},
	"love_you_gesture": Emoji{
		IconURL: fmt.Sprintf("%s/1f91f.png?v8", BaseURL),
		Code:    129311,
	},
	"low_brightness": Emoji{
		IconURL: fmt.Sprintf("%s/1f505.png?v8", BaseURL),
		Code:    128261,
	},
	"luggage": Emoji{
		IconURL: fmt.Sprintf("%s/1f9f3.png?v8", BaseURL),
		Code:    129523,
	},
	"lungs": Emoji{
		IconURL: fmt.Sprintf("%s/1fac1.png?v8", BaseURL),
		Code:    129729,
	},
	"luxembourg": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f1-1f1fa.png?v8", BaseURL),
		Code:    127473,
	},
	"lying_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f925.png?v8", BaseURL),
		Code:    129317,
	},
	"m": Emoji{
		IconURL: fmt.Sprintf("%s/24c2.png?v8", BaseURL),
		Code:    9410,
	},
	"macau": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1f4.png?v8", BaseURL),
		Code:    127474,
	},
	"macedonia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1f0.png?v8", BaseURL),
		Code:    127474,
	},
	"madagascar": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1ec.png?v8", BaseURL),
		Code:    127474,
	},
	"mag": Emoji{
		IconURL: fmt.Sprintf("%s/1f50d.png?v8", BaseURL),
		Code:    128269,
	},
	"mag_right": Emoji{
		IconURL: fmt.Sprintf("%s/1f50e.png?v8", BaseURL),
		Code:    128270,
	},
	"mage": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d9.png?v8", BaseURL),
		Code:    129497,
	},
	"mage_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d9-2642.png?v8", BaseURL),
		Code:    129497,
	},
	"mage_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d9-2640.png?v8", BaseURL),
		Code:    129497,
	},
	"magic_wand": Emoji{
		IconURL: fmt.Sprintf("%s/1fa84.png?v8", BaseURL),
		Code:    129668,
	},
	"magnet": Emoji{
		IconURL: fmt.Sprintf("%s/1f9f2.png?v8", BaseURL),
		Code:    129522,
	},
	"mahjong": Emoji{
		IconURL: fmt.Sprintf("%s/1f004.png?v8", BaseURL),
		Code:    126980,
	},
	"mailbox": Emoji{
		IconURL: fmt.Sprintf("%s/1f4eb.png?v8", BaseURL),
		Code:    128235,
	},
	"mailbox_closed": Emoji{
		IconURL: fmt.Sprintf("%s/1f4ea.png?v8", BaseURL),
		Code:    128234,
	},
	"mailbox_with_mail": Emoji{
		IconURL: fmt.Sprintf("%s/1f4ec.png?v8", BaseURL),
		Code:    128236,
	},
	"mailbox_with_no_mail": Emoji{
		IconURL: fmt.Sprintf("%s/1f4ed.png?v8", BaseURL),
		Code:    128237,
	},
	"malawi": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1fc.png?v8", BaseURL),
		Code:    127474,
	},
	"malaysia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1fe.png?v8", BaseURL),
		Code:    127474,
	},
	"maldives": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1fb.png?v8", BaseURL),
		Code:    127474,
	},
	"male_detective": Emoji{
		IconURL: fmt.Sprintf("%s/1f575-2642.png?v8", BaseURL),
		Code:    128373,
	},
	"male_sign": Emoji{
		IconURL: fmt.Sprintf("%s/2642.png?v8", BaseURL),
		Code:    9794,
	},
	"mali": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1f1.png?v8", BaseURL),
		Code:    127474,
	},
	"malta": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1f9.png?v8", BaseURL),
		Code:    127474,
	},
	"mammoth": Emoji{
		IconURL: fmt.Sprintf("%s/1f9a3.png?v8", BaseURL),
		Code:    129443,
	},
	"man": Emoji{
		IconURL: fmt.Sprintf("%s/1f468.png?v8", BaseURL),
		Code:    128104,
	},
	"man_artist": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f3a8.png?v8", BaseURL),
		Code:    128104,
	},
	"man_astronaut": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f680.png?v8", BaseURL),
		Code:    128104,
	},
	"man_beard": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d4-2642.png?v8", BaseURL),
		Code:    129492,
	},
	"man_cartwheeling": Emoji{
		IconURL: fmt.Sprintf("%s/1f938-2642.png?v8", BaseURL),
		Code:    129336,
	},
	"man_cook": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f373.png?v8", BaseURL),
		Code:    128104,
	},
	"man_dancing": Emoji{
		IconURL: fmt.Sprintf("%s/1f57a.png?v8", BaseURL),
		Code:    128378,
	},
	"man_facepalming": Emoji{
		IconURL: fmt.Sprintf("%s/1f926-2642.png?v8", BaseURL),
		Code:    129318,
	},
	"man_factory_worker": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f3ed.png?v8", BaseURL),
		Code:    128104,
	},
	"man_farmer": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f33e.png?v8", BaseURL),
		Code:    128104,
	},
	"man_feeding_baby": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f37c.png?v8", BaseURL),
		Code:    128104,
	},
	"man_firefighter": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f692.png?v8", BaseURL),
		Code:    128104,
	},
	"man_health_worker": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-2695.png?v8", BaseURL),
		Code:    128104,
	},
	"man_in_manual_wheelchair": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f9bd.png?v8", BaseURL),
		Code:    128104,
	},
	"man_in_motorized_wheelchair": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f9bc.png?v8", BaseURL),
		Code:    128104,
	},
	"man_in_tuxedo": Emoji{
		IconURL: fmt.Sprintf("%s/1f935-2642.png?v8", BaseURL),
		Code:    129333,
	},
	"man_judge": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-2696.png?v8", BaseURL),
		Code:    128104,
	},
	"man_juggling": Emoji{
		IconURL: fmt.Sprintf("%s/1f939-2642.png?v8", BaseURL),
		Code:    129337,
	},
	"man_mechanic": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f527.png?v8", BaseURL),
		Code:    128104,
	},
	"man_office_worker": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f4bc.png?v8", BaseURL),
		Code:    128104,
	},
	"man_pilot": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-2708.png?v8", BaseURL),
		Code:    128104,
	},
	"man_playing_handball": Emoji{
		IconURL: fmt.Sprintf("%s/1f93e-2642.png?v8", BaseURL),
		Code:    129342,
	},
	"man_playing_water_polo": Emoji{
		IconURL: fmt.Sprintf("%s/1f93d-2642.png?v8", BaseURL),
		Code:    129341,
	},
	"man_scientist": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f52c.png?v8", BaseURL),
		Code:    128104,
	},
	"man_shrugging": Emoji{
		IconURL: fmt.Sprintf("%s/1f937-2642.png?v8", BaseURL),
		Code:    129335,
	},
	"man_singer": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f3a4.png?v8", BaseURL),
		Code:    128104,
	},
	"man_student": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f393.png?v8", BaseURL),
		Code:    128104,
	},
	"man_teacher": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f3eb.png?v8", BaseURL),
		Code:    128104,
	},
	"man_technologist": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f4bb.png?v8", BaseURL),
		Code:    128104,
	},
	"man_with_gua_pi_mao": Emoji{
		IconURL: fmt.Sprintf("%s/1f472.png?v8", BaseURL),
		Code:    128114,
	},
	"man_with_probing_cane": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f9af.png?v8", BaseURL),
		Code:    128104,
	},
	"man_with_turban": Emoji{
		IconURL: fmt.Sprintf("%s/1f473-2642.png?v8", BaseURL),
		Code:    128115,
	},
	"man_with_veil": Emoji{
		IconURL: fmt.Sprintf("%s/1f470-2642.png?v8", BaseURL),
		Code:    128112,
	},
	"mandarin": Emoji{
		IconURL: fmt.Sprintf("%s/1f34a.png?v8", BaseURL),
		Code:    127818,
	},
	"mango": Emoji{
		IconURL: fmt.Sprintf("%s/1f96d.png?v8", BaseURL),
		Code:    129389,
	},
	"mans_shoe": Emoji{
		IconURL: fmt.Sprintf("%s/1f45e.png?v8", BaseURL),
		Code:    128094,
	},
	"mantelpiece_clock": Emoji{
		IconURL: fmt.Sprintf("%s/1f570.png?v8", BaseURL),
		Code:    128368,
	},
	"manual_wheelchair": Emoji{
		IconURL: fmt.Sprintf("%s/1f9bd.png?v8", BaseURL),
		Code:    129469,
	},
	"maple_leaf": Emoji{
		IconURL: fmt.Sprintf("%s/1f341.png?v8", BaseURL),
		Code:    127809,
	},
	"marshall_islands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1ed.png?v8", BaseURL),
		Code:    127474,
	},
	"martial_arts_uniform": Emoji{
		IconURL: fmt.Sprintf("%s/1f94b.png?v8", BaseURL),
		Code:    129355,
	},
	"martinique": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1f6.png?v8", BaseURL),
		Code:    127474,
	},
	"mask": Emoji{
		IconURL: fmt.Sprintf("%s/1f637.png?v8", BaseURL),
		Code:    128567,
	},
	"massage": Emoji{
		IconURL: fmt.Sprintf("%s/1f486.png?v8", BaseURL),
		Code:    128134,
	},
	"massage_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f486-2642.png?v8", BaseURL),
		Code:    128134,
	},
	"massage_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f486-2640.png?v8", BaseURL),
		Code:    128134,
	},
	"mate": Emoji{
		IconURL: fmt.Sprintf("%s/1f9c9.png?v8", BaseURL),
		Code:    129481,
	},
	"mauritania": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1f7.png?v8", BaseURL),
		Code:    127474,
	},
	"mauritius": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1fa.png?v8", BaseURL),
		Code:    127474,
	},
	"mayotte": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fe-1f1f9.png?v8", BaseURL),
		Code:    127486,
	},
	"meat_on_bone": Emoji{
		IconURL: fmt.Sprintf("%s/1f356.png?v8", BaseURL),
		Code:    127830,
	},
	"mechanic": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f527.png?v8", BaseURL),
		Code:    129489,
	},
	"mechanical_arm": Emoji{
		IconURL: fmt.Sprintf("%s/1f9be.png?v8", BaseURL),
		Code:    129470,
	},
	"mechanical_leg": Emoji{
		IconURL: fmt.Sprintf("%s/1f9bf.png?v8", BaseURL),
		Code:    129471,
	},
	"medal_military": Emoji{
		IconURL: fmt.Sprintf("%s/1f396.png?v8", BaseURL),
		Code:    127894,
	},
	"medal_sports": Emoji{
		IconURL: fmt.Sprintf("%s/1f3c5.png?v8", BaseURL),
		Code:    127941,
	},
	"medical_symbol": Emoji{
		IconURL: fmt.Sprintf("%s/2695.png?v8", BaseURL),
		Code:    9877,
	},
	"mega": Emoji{
		IconURL: fmt.Sprintf("%s/1f4e3.png?v8", BaseURL),
		Code:    128227,
	},
	"melon": Emoji{
		IconURL: fmt.Sprintf("%s/1f348.png?v8", BaseURL),
		Code:    127816,
	},
	"memo": Emoji{
		IconURL: fmt.Sprintf("%s/1f4dd.png?v8", BaseURL),
		Code:    128221,
	},
	"men_wrestling": Emoji{
		IconURL: fmt.Sprintf("%s/1f93c-2642.png?v8", BaseURL),
		Code:    129340,
	},
	"mending_heart": Emoji{
		IconURL: fmt.Sprintf("%s/2764-1fa79.png?v8", BaseURL),
		Code:    10084,
	},
	"menorah": Emoji{
		IconURL: fmt.Sprintf("%s/1f54e.png?v8", BaseURL),
		Code:    128334,
	},
	"mens": Emoji{
		IconURL: fmt.Sprintf("%s/1f6b9.png?v8", BaseURL),
		Code:    128697,
	},
	"mermaid": Emoji{
		IconURL: fmt.Sprintf("%s/1f9dc-2640.png?v8", BaseURL),
		Code:    129500,
	},
	"merman": Emoji{
		IconURL: fmt.Sprintf("%s/1f9dc-2642.png?v8", BaseURL),
		Code:    129500,
	},
	"merperson": Emoji{
		IconURL: fmt.Sprintf("%s/1f9dc.png?v8", BaseURL),
		Code:    129500,
	},
	"metal": Emoji{
		IconURL: fmt.Sprintf("%s/1f918.png?v8", BaseURL),
		Code:    129304,
	},
	"metro": Emoji{
		IconURL: fmt.Sprintf("%s/1f687.png?v8", BaseURL),
		Code:    128647,
	},
	"mexico": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1fd.png?v8", BaseURL),
		Code:    127474,
	},
	"microbe": Emoji{
		IconURL: fmt.Sprintf("%s/1f9a0.png?v8", BaseURL),
		Code:    129440,
	},
	"micronesia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1eb-1f1f2.png?v8", BaseURL),
		Code:    127467,
	},
	"microphone": Emoji{
		IconURL: fmt.Sprintf("%s/1f3a4.png?v8", BaseURL),
		Code:    127908,
	},
	"microscope": Emoji{
		IconURL: fmt.Sprintf("%s/1f52c.png?v8", BaseURL),
		Code:    128300,
	},
	"middle_finger": Emoji{
		IconURL: fmt.Sprintf("%s/1f595.png?v8", BaseURL),
		Code:    128405,
	},
	"military_helmet": Emoji{
		IconURL: fmt.Sprintf("%s/1fa96.png?v8", BaseURL),
		Code:    129686,
	},
	"milk_glass": Emoji{
		IconURL: fmt.Sprintf("%s/1f95b.png?v8", BaseURL),
		Code:    129371,
	},
	"milky_way": Emoji{
		IconURL: fmt.Sprintf("%s/1f30c.png?v8", BaseURL),
		Code:    127756,
	},
	"minibus": Emoji{
		IconURL: fmt.Sprintf("%s/1f690.png?v8", BaseURL),
		Code:    128656,
	},
	"minidisc": Emoji{
		IconURL: fmt.Sprintf("%s/1f4bd.png?v8", BaseURL),
		Code:    128189,
	},
	"mirror": Emoji{
		IconURL: fmt.Sprintf("%s/1fa9e.png?v8", BaseURL),
		Code:    129694,
	},
	"mobile_phone_off": Emoji{
		IconURL: fmt.Sprintf("%s/1f4f4.png?v8", BaseURL),
		Code:    128244,
	},
	"moldova": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1e9.png?v8", BaseURL),
		Code:    127474,
	},
	"monaco": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1e8.png?v8", BaseURL),
		Code:    127474,
	},
	"money_mouth_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f911.png?v8", BaseURL),
		Code:    129297,
	},
	"money_with_wings": Emoji{
		IconURL: fmt.Sprintf("%s/1f4b8.png?v8", BaseURL),
		Code:    128184,
	},
	"moneybag": Emoji{
		IconURL: fmt.Sprintf("%s/1f4b0.png?v8", BaseURL),
		Code:    128176,
	},
	"mongolia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1f3.png?v8", BaseURL),
		Code:    127474,
	},
	"monkey": Emoji{
		IconURL: fmt.Sprintf("%s/1f412.png?v8", BaseURL),
		Code:    128018,
	},
	"monkey_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f435.png?v8", BaseURL),
		Code:    128053,
	},
	"monocle_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d0.png?v8", BaseURL),
		Code:    129488,
	},
	"monorail": Emoji{
		IconURL: fmt.Sprintf("%s/1f69d.png?v8", BaseURL),
		Code:    128669,
	},
	"montenegro": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1ea.png?v8", BaseURL),
		Code:    127474,
	},
	"montserrat": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1f8.png?v8", BaseURL),
		Code:    127474,
	},
	"moon": Emoji{
		IconURL: fmt.Sprintf("%s/1f314.png?v8", BaseURL),
		Code:    127764,
	},
	"moon_cake": Emoji{
		IconURL: fmt.Sprintf("%s/1f96e.png?v8", BaseURL),
		Code:    129390,
	},
	"morocco": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1e6.png?v8", BaseURL),
		Code:    127474,
	},
	"mortar_board": Emoji{
		IconURL: fmt.Sprintf("%s/1f393.png?v8", BaseURL),
		Code:    127891,
	},
	"mosque": Emoji{
		IconURL: fmt.Sprintf("%s/1f54c.png?v8", BaseURL),
		Code:    128332,
	},
	"mosquito": Emoji{
		IconURL: fmt.Sprintf("%s/1f99f.png?v8", BaseURL),
		Code:    129439,
	},
	"motor_boat": Emoji{
		IconURL: fmt.Sprintf("%s/1f6e5.png?v8", BaseURL),
		Code:    128741,
	},
	"motor_scooter": Emoji{
		IconURL: fmt.Sprintf("%s/1f6f5.png?v8", BaseURL),
		Code:    128757,
	},
	"motorcycle": Emoji{
		IconURL: fmt.Sprintf("%s/1f3cd.png?v8", BaseURL),
		Code:    127949,
	},
	"motorized_wheelchair": Emoji{
		IconURL: fmt.Sprintf("%s/1f9bc.png?v8", BaseURL),
		Code:    129468,
	},
	"motorway": Emoji{
		IconURL: fmt.Sprintf("%s/1f6e3.png?v8", BaseURL),
		Code:    128739,
	},
	"mount_fuji": Emoji{
		IconURL: fmt.Sprintf("%s/1f5fb.png?v8", BaseURL),
		Code:    128507,
	},
	"mountain": Emoji{
		IconURL: fmt.Sprintf("%s/26f0.png?v8", BaseURL),
		Code:    9968,
	},
	"mountain_bicyclist": Emoji{
		IconURL: fmt.Sprintf("%s/1f6b5.png?v8", BaseURL),
		Code:    128693,
	},
	"mountain_biking_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f6b5-2642.png?v8", BaseURL),
		Code:    128693,
	},
	"mountain_biking_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f6b5-2640.png?v8", BaseURL),
		Code:    128693,
	},
	"mountain_cableway": Emoji{
		IconURL: fmt.Sprintf("%s/1f6a0.png?v8", BaseURL),
		Code:    128672,
	},
	"mountain_railway": Emoji{
		IconURL: fmt.Sprintf("%s/1f69e.png?v8", BaseURL),
		Code:    128670,
	},
	"mountain_snow": Emoji{
		IconURL: fmt.Sprintf("%s/1f3d4.png?v8", BaseURL),
		Code:    127956,
	},
	"mouse": Emoji{
		IconURL: fmt.Sprintf("%s/1f42d.png?v8", BaseURL),
		Code:    128045,
	},
	"mouse2": Emoji{
		IconURL: fmt.Sprintf("%s/1f401.png?v8", BaseURL),
		Code:    128001,
	},
	"mouse_trap": Emoji{
		IconURL: fmt.Sprintf("%s/1faa4.png?v8", BaseURL),
		Code:    129700,
	},
	"movie_camera": Emoji{
		IconURL: fmt.Sprintf("%s/1f3a5.png?v8", BaseURL),
		Code:    127909,
	},
	"moyai": Emoji{
		IconURL: fmt.Sprintf("%s/1f5ff.png?v8", BaseURL),
		Code:    128511,
	},
	"mozambique": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1ff.png?v8", BaseURL),
		Code:    127474,
	},
	"mrs_claus": Emoji{
		IconURL: fmt.Sprintf("%s/1f936.png?v8", BaseURL),
		Code:    129334,
	},
	"muscle": Emoji{
		IconURL: fmt.Sprintf("%s/1f4aa.png?v8", BaseURL),
		Code:    128170,
	},
	"mushroom": Emoji{
		IconURL: fmt.Sprintf("%s/1f344.png?v8", BaseURL),
		Code:    127812,
	},
	"musical_keyboard": Emoji{
		IconURL: fmt.Sprintf("%s/1f3b9.png?v8", BaseURL),
		Code:    127929,
	},
	"musical_note": Emoji{
		IconURL: fmt.Sprintf("%s/1f3b5.png?v8", BaseURL),
		Code:    127925,
	},
	"musical_score": Emoji{
		IconURL: fmt.Sprintf("%s/1f3bc.png?v8", BaseURL),
		Code:    127932,
	},
	"mute": Emoji{
		IconURL: fmt.Sprintf("%s/1f507.png?v8", BaseURL),
		Code:    128263,
	},
	"mx_claus": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f384.png?v8", BaseURL),
		Code:    129489,
	},
	"myanmar": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1f2.png?v8", BaseURL),
		Code:    127474,
	},
	"nail_care": Emoji{
		IconURL: fmt.Sprintf("%s/1f485.png?v8", BaseURL),
		Code:    128133,
	},
	"name_badge": Emoji{
		IconURL: fmt.Sprintf("%s/1f4db.png?v8", BaseURL),
		Code:    128219,
	},
	"namibia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f3-1f1e6.png?v8", BaseURL),
		Code:    127475,
	},
	"national_park": Emoji{
		IconURL: fmt.Sprintf("%s/1f3de.png?v8", BaseURL),
		Code:    127966,
	},
	"nauru": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f3-1f1f7.png?v8", BaseURL),
		Code:    127475,
	},
	"nauseated_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f922.png?v8", BaseURL),
		Code:    129314,
	},
	"nazar_amulet": Emoji{
		IconURL: fmt.Sprintf("%s/1f9ff.png?v8", BaseURL),
		Code:    129535,
	},
	"neckbeard": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/neckbeard.png?v8",
		Code:    0,
	},
	"necktie": Emoji{
		IconURL: fmt.Sprintf("%s/1f454.png?v8", BaseURL),
		Code:    128084,
	},
	"negative_squared_cross_mark": Emoji{
		IconURL: fmt.Sprintf("%s/274e.png?v8", BaseURL),
		Code:    10062,
	},
	"nepal": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f3-1f1f5.png?v8", BaseURL),
		Code:    127475,
	},
	"nerd_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f913.png?v8", BaseURL),
		Code:    129299,
	},
	"nesting_dolls": Emoji{
		IconURL: fmt.Sprintf("%s/1fa86.png?v8", BaseURL),
		Code:    129670,
	},
	"netherlands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f3-1f1f1.png?v8", BaseURL),
		Code:    127475,
	},
	"neutral_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f610.png?v8", BaseURL),
		Code:    128528,
	},
	"new": Emoji{
		IconURL: fmt.Sprintf("%s/1f195.png?v8", BaseURL),
		Code:    127381,
	},
	"new_caledonia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f3-1f1e8.png?v8", BaseURL),
		Code:    127475,
	},
	"new_moon": Emoji{
		IconURL: fmt.Sprintf("%s/1f311.png?v8", BaseURL),
		Code:    127761,
	},
	"new_moon_with_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f31a.png?v8", BaseURL),
		Code:    127770,
	},
	"new_zealand": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f3-1f1ff.png?v8", BaseURL),
		Code:    127475,
	},
	"newspaper": Emoji{
		IconURL: fmt.Sprintf("%s/1f4f0.png?v8", BaseURL),
		Code:    128240,
	},
	"newspaper_roll": Emoji{
		IconURL: fmt.Sprintf("%s/1f5de.png?v8", BaseURL),
		Code:    128478,
	},
	"next_track_button": Emoji{
		IconURL: fmt.Sprintf("%s/23ed.png?v8", BaseURL),
		Code:    9197,
	},
	"ng": Emoji{
		IconURL: fmt.Sprintf("%s/1f196.png?v8", BaseURL),
		Code:    127382,
	},
	"ng_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f645-2642.png?v8", BaseURL),
		Code:    128581,
	},
	"ng_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f645-2640.png?v8", BaseURL),
		Code:    128581,
	},
	"nicaragua": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f3-1f1ee.png?v8", BaseURL),
		Code:    127475,
	},
	"niger": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f3-1f1ea.png?v8", BaseURL),
		Code:    127475,
	},
	"nigeria": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f3-1f1ec.png?v8", BaseURL),
		Code:    127475,
	},
	"night_with_stars": Emoji{
		IconURL: fmt.Sprintf("%s/1f303.png?v8", BaseURL),
		Code:    127747,
	},
	"nine": Emoji{
		IconURL: fmt.Sprintf("%s/0039-20e3.png?v8", BaseURL),
		Code:    57,
	},
	"ninja": Emoji{
		IconURL: fmt.Sprintf("%s/1f977.png?v8", BaseURL),
		Code:    129399,
	},
	"niue": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f3-1f1fa.png?v8", BaseURL),
		Code:    127475,
	},
	"no_bell": Emoji{
		IconURL: fmt.Sprintf("%s/1f515.png?v8", BaseURL),
		Code:    128277,
	},
	"no_bicycles": Emoji{
		IconURL: fmt.Sprintf("%s/1f6b3.png?v8", BaseURL),
		Code:    128691,
	},
	"no_entry": Emoji{
		IconURL: fmt.Sprintf("%s/26d4.png?v8", BaseURL),
		Code:    9940,
	},
	"no_entry_sign": Emoji{
		IconURL: fmt.Sprintf("%s/1f6ab.png?v8", BaseURL),
		Code:    128683,
	},
	"no_good": Emoji{
		IconURL: fmt.Sprintf("%s/1f645.png?v8", BaseURL),
		Code:    128581,
	},
	"no_good_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f645-2642.png?v8", BaseURL),
		Code:    128581,
	},
	"no_good_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f645-2640.png?v8", BaseURL),
		Code:    128581,
	},
	"no_mobile_phones": Emoji{
		IconURL: fmt.Sprintf("%s/1f4f5.png?v8", BaseURL),
		Code:    128245,
	},
	"no_mouth": Emoji{
		IconURL: fmt.Sprintf("%s/1f636.png?v8", BaseURL),
		Code:    128566,
	},
	"no_pedestrians": Emoji{
		IconURL: fmt.Sprintf("%s/1f6b7.png?v8", BaseURL),
		Code:    128695,
	},
	"no_smoking": Emoji{
		IconURL: fmt.Sprintf("%s/1f6ad.png?v8", BaseURL),
		Code:    128685,
	},
	"non-potable_water": Emoji{
		IconURL: fmt.Sprintf("%s/1f6b1.png?v8", BaseURL),
		Code:    128689,
	},
	"norfolk_island": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f3-1f1eb.png?v8", BaseURL),
		Code:    127475,
	},
	"north_korea": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f0-1f1f5.png?v8", BaseURL),
		Code:    127472,
	},
	"northern_mariana_islands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1f5.png?v8", BaseURL),
		Code:    127474,
	},
	"norway": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f3-1f1f4.png?v8", BaseURL),
		Code:    127475,
	},
	"nose": Emoji{
		IconURL: fmt.Sprintf("%s/1f443.png?v8", BaseURL),
		Code:    128067,
	},
	"notebook": Emoji{
		IconURL: fmt.Sprintf("%s/1f4d3.png?v8", BaseURL),
		Code:    128211,
	},
	"notebook_with_decorative_cover": Emoji{
		IconURL: fmt.Sprintf("%s/1f4d4.png?v8", BaseURL),
		Code:    128212,
	},
	"notes": Emoji{
		IconURL: fmt.Sprintf("%s/1f3b6.png?v8", BaseURL),
		Code:    127926,
	},
	"nut_and_bolt": Emoji{
		IconURL: fmt.Sprintf("%s/1f529.png?v8", BaseURL),
		Code:    128297,
	},
	"o": Emoji{
		IconURL: fmt.Sprintf("%s/2b55.png?v8", BaseURL),
		Code:    11093,
	},
	"o2": Emoji{
		IconURL: fmt.Sprintf("%s/1f17e.png?v8", BaseURL),
		Code:    127358,
	},
	"ocean": Emoji{
		IconURL: fmt.Sprintf("%s/1f30a.png?v8", BaseURL),
		Code:    127754,
	},
	"octocat": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/octocat.png?v8",
		Code:    0,
	},
	"octopus": Emoji{
		IconURL: fmt.Sprintf("%s/1f419.png?v8", BaseURL),
		Code:    128025,
	},
	"oden": Emoji{
		IconURL: fmt.Sprintf("%s/1f362.png?v8", BaseURL),
		Code:    127842,
	},
	"office": Emoji{
		IconURL: fmt.Sprintf("%s/1f3e2.png?v8", BaseURL),
		Code:    127970,
	},
	"office_worker": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f4bc.png?v8", BaseURL),
		Code:    129489,
	},
	"oil_drum": Emoji{
		IconURL: fmt.Sprintf("%s/1f6e2.png?v8", BaseURL),
		Code:    128738,
	},
	"ok": Emoji{
		IconURL: fmt.Sprintf("%s/1f197.png?v8", BaseURL),
		Code:    127383,
	},
	"ok_hand": Emoji{
		IconURL: fmt.Sprintf("%s/1f44c.png?v8", BaseURL),
		Code:    128076,
	},
	"ok_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f646-2642.png?v8", BaseURL),
		Code:    128582,
	},
	"ok_person": Emoji{
		IconURL: fmt.Sprintf("%s/1f646.png?v8", BaseURL),
		Code:    128582,
	},
	"ok_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f646-2640.png?v8", BaseURL),
		Code:    128582,
	},
	"old_key": Emoji{
		IconURL: fmt.Sprintf("%s/1f5dd.png?v8", BaseURL),
		Code:    128477,
	},
	"older_adult": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d3.png?v8", BaseURL),
		Code:    129491,
	},
	"older_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f474.png?v8", BaseURL),
		Code:    128116,
	},
	"older_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f475.png?v8", BaseURL),
		Code:    128117,
	},
	"olive": Emoji{
		IconURL: fmt.Sprintf("%s/1fad2.png?v8", BaseURL),
		Code:    129746,
	},
	"om": Emoji{
		IconURL: fmt.Sprintf("%s/1f549.png?v8", BaseURL),
		Code:    128329,
	},
	"oman": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f4-1f1f2.png?v8", BaseURL),
		Code:    127476,
	},
	"on": Emoji{
		IconURL: fmt.Sprintf("%s/1f51b.png?v8", BaseURL),
		Code:    128283,
	},
	"oncoming_automobile": Emoji{
		IconURL: fmt.Sprintf("%s/1f698.png?v8", BaseURL),
		Code:    128664,
	},
	"oncoming_bus": Emoji{
		IconURL: fmt.Sprintf("%s/1f68d.png?v8", BaseURL),
		Code:    128653,
	},
	"oncoming_police_car": Emoji{
		IconURL: fmt.Sprintf("%s/1f694.png?v8", BaseURL),
		Code:    128660,
	},
	"oncoming_taxi": Emoji{
		IconURL: fmt.Sprintf("%s/1f696.png?v8", BaseURL),
		Code:    128662,
	},
	"one": Emoji{
		IconURL: fmt.Sprintf("%s/0031-20e3.png?v8", BaseURL),
		Code:    49,
	},
	"one_piece_swimsuit": Emoji{
		IconURL: fmt.Sprintf("%s/1fa71.png?v8", BaseURL),
		Code:    129649,
	},
	"onion": Emoji{
		IconURL: fmt.Sprintf("%s/1f9c5.png?v8", BaseURL),
		Code:    129477,
	},
	"open_book": Emoji{
		IconURL: fmt.Sprintf("%s/1f4d6.png?v8", BaseURL),
		Code:    128214,
	},
	"open_file_folder": Emoji{
		IconURL: fmt.Sprintf("%s/1f4c2.png?v8", BaseURL),
		Code:    128194,
	},
	"open_hands": Emoji{
		IconURL: fmt.Sprintf("%s/1f450.png?v8", BaseURL),
		Code:    128080,
	},
	"open_mouth": Emoji{
		IconURL: fmt.Sprintf("%s/1f62e.png?v8", BaseURL),
		Code:    128558,
	},
	"open_umbrella": Emoji{
		IconURL: fmt.Sprintf("%s/2602.png?v8", BaseURL),
		Code:    9730,
	},
	"ophiuchus": Emoji{
		IconURL: fmt.Sprintf("%s/26ce.png?v8", BaseURL),
		Code:    9934,
	},
	"orange": Emoji{
		IconURL: fmt.Sprintf("%s/1f34a.png?v8", BaseURL),
		Code:    127818,
	},
	"orange_book": Emoji{
		IconURL: fmt.Sprintf("%s/1f4d9.png?v8", BaseURL),
		Code:    128217,
	},
	"orange_circle": Emoji{
		IconURL: fmt.Sprintf("%s/1f7e0.png?v8", BaseURL),
		Code:    128992,
	},
	"orange_heart": Emoji{
		IconURL: fmt.Sprintf("%s/1f9e1.png?v8", BaseURL),
		Code:    129505,
	},
	"orange_square": Emoji{
		IconURL: fmt.Sprintf("%s/1f7e7.png?v8", BaseURL),
		Code:    128999,
	},
	"orangutan": Emoji{
		IconURL: fmt.Sprintf("%s/1f9a7.png?v8", BaseURL),
		Code:    129447,
	},
	"orthodox_cross": Emoji{
		IconURL: fmt.Sprintf("%s/2626.png?v8", BaseURL),
		Code:    9766,
	},
	"otter": Emoji{
		IconURL: fmt.Sprintf("%s/1f9a6.png?v8", BaseURL),
		Code:    129446,
	},
	"outbox_tray": Emoji{
		IconURL: fmt.Sprintf("%s/1f4e4.png?v8", BaseURL),
		Code:    128228,
	},
	"owl": Emoji{
		IconURL: fmt.Sprintf("%s/1f989.png?v8", BaseURL),
		Code:    129417,
	},
	"ox": Emoji{
		IconURL: fmt.Sprintf("%s/1f402.png?v8", BaseURL),
		Code:    128002,
	},
	"oyster": Emoji{
		IconURL: fmt.Sprintf("%s/1f9aa.png?v8", BaseURL),
		Code:    129450,
	},
	"package": Emoji{
		IconURL: fmt.Sprintf("%s/1f4e6.png?v8", BaseURL),
		Code:    128230,
	},
	"page_facing_up": Emoji{
		IconURL: fmt.Sprintf("%s/1f4c4.png?v8", BaseURL),
		Code:    128196,
	},
	"page_with_curl": Emoji{
		IconURL: fmt.Sprintf("%s/1f4c3.png?v8", BaseURL),
		Code:    128195,
	},
	"pager": Emoji{
		IconURL: fmt.Sprintf("%s/1f4df.png?v8", BaseURL),
		Code:    128223,
	},
	"paintbrush": Emoji{
		IconURL: fmt.Sprintf("%s/1f58c.png?v8", BaseURL),
		Code:    128396,
	},
	"pakistan": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f5-1f1f0.png?v8", BaseURL),
		Code:    127477,
	},
	"palau": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f5-1f1fc.png?v8", BaseURL),
		Code:    127477,
	},
	"palestinian_territories": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f5-1f1f8.png?v8", BaseURL),
		Code:    127477,
	},
	"palm_tree": Emoji{
		IconURL: fmt.Sprintf("%s/1f334.png?v8", BaseURL),
		Code:    127796,
	},
	"palms_up_together": Emoji{
		IconURL: fmt.Sprintf("%s/1f932.png?v8", BaseURL),
		Code:    129330,
	},
	"panama": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f5-1f1e6.png?v8", BaseURL),
		Code:    127477,
	},
	"pancakes": Emoji{
		IconURL: fmt.Sprintf("%s/1f95e.png?v8", BaseURL),
		Code:    129374,
	},
	"panda_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f43c.png?v8", BaseURL),
		Code:    128060,
	},
	"paperclip": Emoji{
		IconURL: fmt.Sprintf("%s/1f4ce.png?v8", BaseURL),
		Code:    128206,
	},
	"paperclips": Emoji{
		IconURL: fmt.Sprintf("%s/1f587.png?v8", BaseURL),
		Code:    128391,
	},
	"papua_new_guinea": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f5-1f1ec.png?v8", BaseURL),
		Code:    127477,
	},
	"parachute": Emoji{
		IconURL: fmt.Sprintf("%s/1fa82.png?v8", BaseURL),
		Code:    129666,
	},
	"paraguay": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f5-1f1fe.png?v8", BaseURL),
		Code:    127477,
	},
	"parasol_on_ground": Emoji{
		IconURL: fmt.Sprintf("%s/26f1.png?v8", BaseURL),
		Code:    9969,
	},
	"parking": Emoji{
		IconURL: fmt.Sprintf("%s/1f17f.png?v8", BaseURL),
		Code:    127359,
	},
	"parrot": Emoji{
		IconURL: fmt.Sprintf("%s/1f99c.png?v8", BaseURL),
		Code:    129436,
	},
	"part_alternation_mark": Emoji{
		IconURL: fmt.Sprintf("%s/303d.png?v8", BaseURL),
		Code:    12349,
	},
	"partly_sunny": Emoji{
		IconURL: fmt.Sprintf("%s/26c5.png?v8", BaseURL),
		Code:    9925,
	},
	"partying_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f973.png?v8", BaseURL),
		Code:    129395,
	},
	"passenger_ship": Emoji{
		IconURL: fmt.Sprintf("%s/1f6f3.png?v8", BaseURL),
		Code:    128755,
	},
	"passport_control": Emoji{
		IconURL: fmt.Sprintf("%s/1f6c2.png?v8", BaseURL),
		Code:    128706,
	},
	"pause_button": Emoji{
		IconURL: fmt.Sprintf("%s/23f8.png?v8", BaseURL),
		Code:    9208,
	},
	"paw_prints": Emoji{
		IconURL: fmt.Sprintf("%s/1f43e.png?v8", BaseURL),
		Code:    128062,
	},
	"peace_symbol": Emoji{
		IconURL: fmt.Sprintf("%s/262e.png?v8", BaseURL),
		Code:    9774,
	},
	"peach": Emoji{
		IconURL: fmt.Sprintf("%s/1f351.png?v8", BaseURL),
		Code:    127825,
	},
	"peacock": Emoji{
		IconURL: fmt.Sprintf("%s/1f99a.png?v8", BaseURL),
		Code:    129434,
	},
	"peanuts": Emoji{
		IconURL: fmt.Sprintf("%s/1f95c.png?v8", BaseURL),
		Code:    129372,
	},
	"pear": Emoji{
		IconURL: fmt.Sprintf("%s/1f350.png?v8", BaseURL),
		Code:    127824,
	},
	"pen": Emoji{
		IconURL: fmt.Sprintf("%s/1f58a.png?v8", BaseURL),
		Code:    128394,
	},
	"pencil": Emoji{
		IconURL: fmt.Sprintf("%s/1f4dd.png?v8", BaseURL),
		Code:    128221,
	},
	"pencil2": Emoji{
		IconURL: fmt.Sprintf("%s/270f.png?v8", BaseURL),
		Code:    9999,
	},
	"penguin": Emoji{
		IconURL: fmt.Sprintf("%s/1f427.png?v8", BaseURL),
		Code:    128039,
	},
	"pensive": Emoji{
		IconURL: fmt.Sprintf("%s/1f614.png?v8", BaseURL),
		Code:    128532,
	},
	"people_holding_hands": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f91d-1f9d1.png?v8", BaseURL),
		Code:    129489,
	},
	"people_hugging": Emoji{
		IconURL: fmt.Sprintf("%s/1fac2.png?v8", BaseURL),
		Code:    129730,
	},
	"performing_arts": Emoji{
		IconURL: fmt.Sprintf("%s/1f3ad.png?v8", BaseURL),
		Code:    127917,
	},
	"persevere": Emoji{
		IconURL: fmt.Sprintf("%s/1f623.png?v8", BaseURL),
		Code:    128547,
	},
	"person_bald": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f9b2.png?v8", BaseURL),
		Code:    129489,
	},
	"person_curly_hair": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f9b1.png?v8", BaseURL),
		Code:    129489,
	},
	"person_feeding_baby": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f37c.png?v8", BaseURL),
		Code:    129489,
	},
	"person_fencing": Emoji{
		IconURL: fmt.Sprintf("%s/1f93a.png?v8", BaseURL),
		Code:    129338,
	},
	"person_in_manual_wheelchair": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f9bd.png?v8", BaseURL),
		Code:    129489,
	},
	"person_in_motorized_wheelchair": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f9bc.png?v8", BaseURL),
		Code:    129489,
	},
	"person_in_tuxedo": Emoji{
		IconURL: fmt.Sprintf("%s/1f935.png?v8", BaseURL),
		Code:    129333,
	},
	"person_red_hair": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f9b0.png?v8", BaseURL),
		Code:    129489,
	},
	"person_white_hair": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f9b3.png?v8", BaseURL),
		Code:    129489,
	},
	"person_with_probing_cane": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f9af.png?v8", BaseURL),
		Code:    129489,
	},
	"person_with_turban": Emoji{
		IconURL: fmt.Sprintf("%s/1f473.png?v8", BaseURL),
		Code:    128115,
	},
	"person_with_veil": Emoji{
		IconURL: fmt.Sprintf("%s/1f470.png?v8", BaseURL),
		Code:    128112,
	},
	"peru": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f5-1f1ea.png?v8", BaseURL),
		Code:    127477,
	},
	"petri_dish": Emoji{
		IconURL: fmt.Sprintf("%s/1f9eb.png?v8", BaseURL),
		Code:    129515,
	},
	"philippines": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f5-1f1ed.png?v8", BaseURL),
		Code:    127477,
	},
	"phone": Emoji{
		IconURL: fmt.Sprintf("%s/260e.png?v8", BaseURL),
		Code:    9742,
	},
	"pick": Emoji{
		IconURL: fmt.Sprintf("%s/26cf.png?v8", BaseURL),
		Code:    9935,
	},
	"pickup_truck": Emoji{
		IconURL: fmt.Sprintf("%s/1f6fb.png?v8", BaseURL),
		Code:    128763,
	},
	"pie": Emoji{
		IconURL: fmt.Sprintf("%s/1f967.png?v8", BaseURL),
		Code:    129383,
	},
	"pig": Emoji{
		IconURL: fmt.Sprintf("%s/1f437.png?v8", BaseURL),
		Code:    128055,
	},
	"pig2": Emoji{
		IconURL: fmt.Sprintf("%s/1f416.png?v8", BaseURL),
		Code:    128022,
	},
	"pig_nose": Emoji{
		IconURL: fmt.Sprintf("%s/1f43d.png?v8", BaseURL),
		Code:    128061,
	},
	"pill": Emoji{
		IconURL: fmt.Sprintf("%s/1f48a.png?v8", BaseURL),
		Code:    128138,
	},
	"pilot": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-2708.png?v8", BaseURL),
		Code:    129489,
	},
	"pinata": Emoji{
		IconURL: fmt.Sprintf("%s/1fa85.png?v8", BaseURL),
		Code:    129669,
	},
	"pinched_fingers": Emoji{
		IconURL: fmt.Sprintf("%s/1f90c.png?v8", BaseURL),
		Code:    129292,
	},
	"pinching_hand": Emoji{
		IconURL: fmt.Sprintf("%s/1f90f.png?v8", BaseURL),
		Code:    129295,
	},
	"pineapple": Emoji{
		IconURL: fmt.Sprintf("%s/1f34d.png?v8", BaseURL),
		Code:    127821,
	},
	"ping_pong": Emoji{
		IconURL: fmt.Sprintf("%s/1f3d3.png?v8", BaseURL),
		Code:    127955,
	},
	"pirate_flag": Emoji{
		IconURL: fmt.Sprintf("%s/1f3f4-2620.png?v8", BaseURL),
		Code:    127988,
	},
	"pisces": Emoji{
		IconURL: fmt.Sprintf("%s/2653.png?v8", BaseURL),
		Code:    9811,
	},
	"pitcairn_islands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f5-1f1f3.png?v8", BaseURL),
		Code:    127477,
	},
	"pizza": Emoji{
		IconURL: fmt.Sprintf("%s/1f355.png?v8", BaseURL),
		Code:    127829,
	},
	"placard": Emoji{
		IconURL: fmt.Sprintf("%s/1faa7.png?v8", BaseURL),
		Code:    129703,
	},
	"place_of_worship": Emoji{
		IconURL: fmt.Sprintf("%s/1f6d0.png?v8", BaseURL),
		Code:    128720,
	},
	"plate_with_cutlery": Emoji{
		IconURL: fmt.Sprintf("%s/1f37d.png?v8", BaseURL),
		Code:    127869,
	},
	"play_or_pause_button": Emoji{
		IconURL: fmt.Sprintf("%s/23ef.png?v8", BaseURL),
		Code:    9199,
	},
	"pleading_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f97a.png?v8", BaseURL),
		Code:    129402,
	},
	"plunger": Emoji{
		IconURL: fmt.Sprintf("%s/1faa0.png?v8", BaseURL),
		Code:    129696,
	},
	"point_down": Emoji{
		IconURL: fmt.Sprintf("%s/1f447.png?v8", BaseURL),
		Code:    128071,
	},
	"point_left": Emoji{
		IconURL: fmt.Sprintf("%s/1f448.png?v8", BaseURL),
		Code:    128072,
	},
	"point_right": Emoji{
		IconURL: fmt.Sprintf("%s/1f449.png?v8", BaseURL),
		Code:    128073,
	},
	"point_up": Emoji{
		IconURL: fmt.Sprintf("%s/261d.png?v8", BaseURL),
		Code:    9757,
	},
	"point_up_2": Emoji{
		IconURL: fmt.Sprintf("%s/1f446.png?v8", BaseURL),
		Code:    128070,
	},
	"poland": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f5-1f1f1.png?v8", BaseURL),
		Code:    127477,
	},
	"polar_bear": Emoji{
		IconURL: fmt.Sprintf("%s/1f43b-2744.png?v8", BaseURL),
		Code:    128059,
	},
	"police_car": Emoji{
		IconURL: fmt.Sprintf("%s/1f693.png?v8", BaseURL),
		Code:    128659,
	},
	"police_officer": Emoji{
		IconURL: fmt.Sprintf("%s/1f46e.png?v8", BaseURL),
		Code:    128110,
	},
	"policeman": Emoji{
		IconURL: fmt.Sprintf("%s/1f46e-2642.png?v8", BaseURL),
		Code:    128110,
	},
	"policewoman": Emoji{
		IconURL: fmt.Sprintf("%s/1f46e-2640.png?v8", BaseURL),
		Code:    128110,
	},
	"poodle": Emoji{
		IconURL: fmt.Sprintf("%s/1f429.png?v8", BaseURL),
		Code:    128041,
	},
	"poop": Emoji{
		IconURL: fmt.Sprintf("%s/1f4a9.png?v8", BaseURL),
		Code:    128169,
	},
	"popcorn": Emoji{
		IconURL: fmt.Sprintf("%s/1f37f.png?v8", BaseURL),
		Code:    127871,
	},
	"portugal": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f5-1f1f9.png?v8", BaseURL),
		Code:    127477,
	},
	"post_office": Emoji{
		IconURL: fmt.Sprintf("%s/1f3e3.png?v8", BaseURL),
		Code:    127971,
	},
	"postal_horn": Emoji{
		IconURL: fmt.Sprintf("%s/1f4ef.png?v8", BaseURL),
		Code:    128239,
	},
	"postbox": Emoji{
		IconURL: fmt.Sprintf("%s/1f4ee.png?v8", BaseURL),
		Code:    128238,
	},
	"potable_water": Emoji{
		IconURL: fmt.Sprintf("%s/1f6b0.png?v8", BaseURL),
		Code:    128688,
	},
	"potato": Emoji{
		IconURL: fmt.Sprintf("%s/1f954.png?v8", BaseURL),
		Code:    129364,
	},
	"potted_plant": Emoji{
		IconURL: fmt.Sprintf("%s/1fab4.png?v8", BaseURL),
		Code:    129716,
	},
	"pouch": Emoji{
		IconURL: fmt.Sprintf("%s/1f45d.png?v8", BaseURL),
		Code:    128093,
	},
	"poultry_leg": Emoji{
		IconURL: fmt.Sprintf("%s/1f357.png?v8", BaseURL),
		Code:    127831,
	},
	"pound": Emoji{
		IconURL: fmt.Sprintf("%s/1f4b7.png?v8", BaseURL),
		Code:    128183,
	},
	"pout": Emoji{
		IconURL: fmt.Sprintf("%s/1f621.png?v8", BaseURL),
		Code:    128545,
	},
	"pouting_cat": Emoji{
		IconURL: fmt.Sprintf("%s/1f63e.png?v8", BaseURL),
		Code:    128574,
	},
	"pouting_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f64e.png?v8", BaseURL),
		Code:    128590,
	},
	"pouting_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f64e-2642.png?v8", BaseURL),
		Code:    128590,
	},
	"pouting_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f64e-2640.png?v8", BaseURL),
		Code:    128590,
	},
	"pray": Emoji{
		IconURL: fmt.Sprintf("%s/1f64f.png?v8", BaseURL),
		Code:    128591,
	},
	"prayer_beads": Emoji{
		IconURL: fmt.Sprintf("%s/1f4ff.png?v8", BaseURL),
		Code:    128255,
	},
	"pregnant_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f930.png?v8", BaseURL),
		Code:    129328,
	},
	"pretzel": Emoji{
		IconURL: fmt.Sprintf("%s/1f968.png?v8", BaseURL),
		Code:    129384,
	},
	"previous_track_button": Emoji{
		IconURL: fmt.Sprintf("%s/23ee.png?v8", BaseURL),
		Code:    9198,
	},
	"prince": Emoji{
		IconURL: fmt.Sprintf("%s/1f934.png?v8", BaseURL),
		Code:    129332,
	},
	"princess": Emoji{
		IconURL: fmt.Sprintf("%s/1f478.png?v8", BaseURL),
		Code:    128120,
	},
	"printer": Emoji{
		IconURL: fmt.Sprintf("%s/1f5a8.png?v8", BaseURL),
		Code:    128424,
	},
	"probing_cane": Emoji{
		IconURL: fmt.Sprintf("%s/1f9af.png?v8", BaseURL),
		Code:    129455,
	},
	"puerto_rico": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f5-1f1f7.png?v8", BaseURL),
		Code:    127477,
	},
	"punch": Emoji{
		IconURL: fmt.Sprintf("%s/1f44a.png?v8", BaseURL),
		Code:    128074,
	},
	"purple_circle": Emoji{
		IconURL: fmt.Sprintf("%s/1f7e3.png?v8", BaseURL),
		Code:    128995,
	},
	"purple_heart": Emoji{
		IconURL: fmt.Sprintf("%s/1f49c.png?v8", BaseURL),
		Code:    128156,
	},
	"purple_square": Emoji{
		IconURL: fmt.Sprintf("%s/1f7ea.png?v8", BaseURL),
		Code:    129002,
	},
	"purse": Emoji{
		IconURL: fmt.Sprintf("%s/1f45b.png?v8", BaseURL),
		Code:    128091,
	},
	"pushpin": Emoji{
		IconURL: fmt.Sprintf("%s/1f4cc.png?v8", BaseURL),
		Code:    128204,
	},
	"put_litter_in_its_place": Emoji{
		IconURL: fmt.Sprintf("%s/1f6ae.png?v8", BaseURL),
		Code:    128686,
	},
	"qatar": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f6-1f1e6.png?v8", BaseURL),
		Code:    127478,
	},
	"question": Emoji{
		IconURL: fmt.Sprintf("%s/2753.png?v8", BaseURL),
		Code:    10067,
	},
	"rabbit": Emoji{
		IconURL: fmt.Sprintf("%s/1f430.png?v8", BaseURL),
		Code:    128048,
	},
	"rabbit2": Emoji{
		IconURL: fmt.Sprintf("%s/1f407.png?v8", BaseURL),
		Code:    128007,
	},
	"raccoon": Emoji{
		IconURL: fmt.Sprintf("%s/1f99d.png?v8", BaseURL),
		Code:    129437,
	},
	"racehorse": Emoji{
		IconURL: fmt.Sprintf("%s/1f40e.png?v8", BaseURL),
		Code:    128014,
	},
	"racing_car": Emoji{
		IconURL: fmt.Sprintf("%s/1f3ce.png?v8", BaseURL),
		Code:    127950,
	},
	"radio": Emoji{
		IconURL: fmt.Sprintf("%s/1f4fb.png?v8", BaseURL),
		Code:    128251,
	},
	"radio_button": Emoji{
		IconURL: fmt.Sprintf("%s/1f518.png?v8", BaseURL),
		Code:    128280,
	},
	"radioactive": Emoji{
		IconURL: fmt.Sprintf("%s/2622.png?v8", BaseURL),
		Code:    9762,
	},
	"rage": Emoji{
		IconURL: fmt.Sprintf("%s/1f621.png?v8", BaseURL),
		Code:    128545,
	},
	"rage1": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/rage1.png?v8",
		Code:    0,
	},
	"rage2": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/rage2.png?v8",
		Code:    0,
	},
	"rage3": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/rage3.png?v8",
		Code:    0,
	},
	"rage4": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/rage4.png?v8",
		Code:    0,
	},
	"railway_car": Emoji{
		IconURL: fmt.Sprintf("%s/1f683.png?v8", BaseURL),
		Code:    128643,
	},
	"railway_track": Emoji{
		IconURL: fmt.Sprintf("%s/1f6e4.png?v8", BaseURL),
		Code:    128740,
	},
	"rainbow": Emoji{
		IconURL: fmt.Sprintf("%s/1f308.png?v8", BaseURL),
		Code:    127752,
	},
	"rainbow_flag": Emoji{
		IconURL: fmt.Sprintf("%s/1f3f3-1f308.png?v8", BaseURL),
		Code:    127987,
	},
	"raised_back_of_hand": Emoji{
		IconURL: fmt.Sprintf("%s/1f91a.png?v8", BaseURL),
		Code:    129306,
	},
	"raised_eyebrow": Emoji{
		IconURL: fmt.Sprintf("%s/1f928.png?v8", BaseURL),
		Code:    129320,
	},
	"raised_hand": Emoji{
		IconURL: fmt.Sprintf("%s/270b.png?v8", BaseURL),
		Code:    9995,
	},
	"raised_hand_with_fingers_splayed": Emoji{
		IconURL: fmt.Sprintf("%s/1f590.png?v8", BaseURL),
		Code:    128400,
	},
	"raised_hands": Emoji{
		IconURL: fmt.Sprintf("%s/1f64c.png?v8", BaseURL),
		Code:    128588,
	},
	"raising_hand": Emoji{
		IconURL: fmt.Sprintf("%s/1f64b.png?v8", BaseURL),
		Code:    128587,
	},
	"raising_hand_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f64b-2642.png?v8", BaseURL),
		Code:    128587,
	},
	"raising_hand_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f64b-2640.png?v8", BaseURL),
		Code:    128587,
	},
	"ram": Emoji{
		IconURL: fmt.Sprintf("%s/1f40f.png?v8", BaseURL),
		Code:    128015,
	},
	"ramen": Emoji{
		IconURL: fmt.Sprintf("%s/1f35c.png?v8", BaseURL),
		Code:    127836,
	},
	"rat": Emoji{
		IconURL: fmt.Sprintf("%s/1f400.png?v8", BaseURL),
		Code:    128000,
	},
	"razor": Emoji{
		IconURL: fmt.Sprintf("%s/1fa92.png?v8", BaseURL),
		Code:    129682,
	},
	"receipt": Emoji{
		IconURL: fmt.Sprintf("%s/1f9fe.png?v8", BaseURL),
		Code:    129534,
	},
	"record_button": Emoji{
		IconURL: fmt.Sprintf("%s/23fa.png?v8", BaseURL),
		Code:    9210,
	},
	"recycle": Emoji{
		IconURL: fmt.Sprintf("%s/267b.png?v8", BaseURL),
		Code:    9851,
	},
	"red_car": Emoji{
		IconURL: fmt.Sprintf("%s/1f697.png?v8", BaseURL),
		Code:    128663,
	},
	"red_circle": Emoji{
		IconURL: fmt.Sprintf("%s/1f534.png?v8", BaseURL),
		Code:    128308,
	},
	"red_envelope": Emoji{
		IconURL: fmt.Sprintf("%s/1f9e7.png?v8", BaseURL),
		Code:    129511,
	},
	"red_haired_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f9b0.png?v8", BaseURL),
		Code:    128104,
	},
	"red_haired_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f9b0.png?v8", BaseURL),
		Code:    128105,
	},
	"red_square": Emoji{
		IconURL: fmt.Sprintf("%s/1f7e5.png?v8", BaseURL),
		Code:    128997,
	},
	"registered": Emoji{
		IconURL: fmt.Sprintf("%s/00ae.png?v8", BaseURL),
		Code:    174,
	},
	"relaxed": Emoji{
		IconURL: fmt.Sprintf("%s/263a.png?v8", BaseURL),
		Code:    9786,
	},
	"relieved": Emoji{
		IconURL: fmt.Sprintf("%s/1f60c.png?v8", BaseURL),
		Code:    128524,
	},
	"reminder_ribbon": Emoji{
		IconURL: fmt.Sprintf("%s/1f397.png?v8", BaseURL),
		Code:    127895,
	},
	"repeat": Emoji{
		IconURL: fmt.Sprintf("%s/1f501.png?v8", BaseURL),
		Code:    128257,
	},
	"repeat_one": Emoji{
		IconURL: fmt.Sprintf("%s/1f502.png?v8", BaseURL),
		Code:    128258,
	},
	"rescue_worker_helmet": Emoji{
		IconURL: fmt.Sprintf("%s/26d1.png?v8", BaseURL),
		Code:    9937,
	},
	"restroom": Emoji{
		IconURL: fmt.Sprintf("%s/1f6bb.png?v8", BaseURL),
		Code:    128699,
	},
	"reunion": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f7-1f1ea.png?v8", BaseURL),
		Code:    127479,
	},
	"revolving_hearts": Emoji{
		IconURL: fmt.Sprintf("%s/1f49e.png?v8", BaseURL),
		Code:    128158,
	},
	"rewind": Emoji{
		IconURL: fmt.Sprintf("%s/23ea.png?v8", BaseURL),
		Code:    9194,
	},
	"rhinoceros": Emoji{
		IconURL: fmt.Sprintf("%s/1f98f.png?v8", BaseURL),
		Code:    129423,
	},
	"ribbon": Emoji{
		IconURL: fmt.Sprintf("%s/1f380.png?v8", BaseURL),
		Code:    127872,
	},
	"rice": Emoji{
		IconURL: fmt.Sprintf("%s/1f35a.png?v8", BaseURL),
		Code:    127834,
	},
	"rice_ball": Emoji{
		IconURL: fmt.Sprintf("%s/1f359.png?v8", BaseURL),
		Code:    127833,
	},
	"rice_cracker": Emoji{
		IconURL: fmt.Sprintf("%s/1f358.png?v8", BaseURL),
		Code:    127832,
	},
	"rice_scene": Emoji{
		IconURL: fmt.Sprintf("%s/1f391.png?v8", BaseURL),
		Code:    127889,
	},
	"right_anger_bubble": Emoji{
		IconURL: fmt.Sprintf("%s/1f5ef.png?v8", BaseURL),
		Code:    128495,
	},
	"ring": Emoji{
		IconURL: fmt.Sprintf("%s/1f48d.png?v8", BaseURL),
		Code:    128141,
	},
	"ringed_planet": Emoji{
		IconURL: fmt.Sprintf("%s/1fa90.png?v8", BaseURL),
		Code:    129680,
	},
	"robot": Emoji{
		IconURL: fmt.Sprintf("%s/1f916.png?v8", BaseURL),
		Code:    129302,
	},
	"rock": Emoji{
		IconURL: fmt.Sprintf("%s/1faa8.png?v8", BaseURL),
		Code:    129704,
	},
	"rocket": Emoji{
		IconURL: fmt.Sprintf("%s/1f680.png?v8", BaseURL),
		Code:    128640,
	},
	"rofl": Emoji{
		IconURL: fmt.Sprintf("%s/1f923.png?v8", BaseURL),
		Code:    129315,
	},
	"roll_eyes": Emoji{
		IconURL: fmt.Sprintf("%s/1f644.png?v8", BaseURL),
		Code:    128580,
	},
	"roll_of_paper": Emoji{
		IconURL: fmt.Sprintf("%s/1f9fb.png?v8", BaseURL),
		Code:    129531,
	},
	"roller_coaster": Emoji{
		IconURL: fmt.Sprintf("%s/1f3a2.png?v8", BaseURL),
		Code:    127906,
	},
	"roller_skate": Emoji{
		IconURL: fmt.Sprintf("%s/1f6fc.png?v8", BaseURL),
		Code:    128764,
	},
	"romania": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f7-1f1f4.png?v8", BaseURL),
		Code:    127479,
	},
	"rooster": Emoji{
		IconURL: fmt.Sprintf("%s/1f413.png?v8", BaseURL),
		Code:    128019,
	},
	"rose": Emoji{
		IconURL: fmt.Sprintf("%s/1f339.png?v8", BaseURL),
		Code:    127801,
	},
	"rosette": Emoji{
		IconURL: fmt.Sprintf("%s/1f3f5.png?v8", BaseURL),
		Code:    127989,
	},
	"rotating_light": Emoji{
		IconURL: fmt.Sprintf("%s/1f6a8.png?v8", BaseURL),
		Code:    128680,
	},
	"round_pushpin": Emoji{
		IconURL: fmt.Sprintf("%s/1f4cd.png?v8", BaseURL),
		Code:    128205,
	},
	"rowboat": Emoji{
		IconURL: fmt.Sprintf("%s/1f6a3.png?v8", BaseURL),
		Code:    128675,
	},
	"rowing_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f6a3-2642.png?v8", BaseURL),
		Code:    128675,
	},
	"rowing_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f6a3-2640.png?v8", BaseURL),
		Code:    128675,
	},
	"ru": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f7-1f1fa.png?v8", BaseURL),
		Code:    127479,
	},
	"rugby_football": Emoji{
		IconURL: fmt.Sprintf("%s/1f3c9.png?v8", BaseURL),
		Code:    127945,
	},
	"runner": Emoji{
		IconURL: fmt.Sprintf("%s/1f3c3.png?v8", BaseURL),
		Code:    127939,
	},
	"running": Emoji{
		IconURL: fmt.Sprintf("%s/1f3c3.png?v8", BaseURL),
		Code:    127939,
	},
	"running_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f3c3-2642.png?v8", BaseURL),
		Code:    127939,
	},
	"running_shirt_with_sash": Emoji{
		IconURL: fmt.Sprintf("%s/1f3bd.png?v8", BaseURL),
		Code:    127933,
	},
	"running_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f3c3-2640.png?v8", BaseURL),
		Code:    127939,
	},
	"rwanda": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f7-1f1fc.png?v8", BaseURL),
		Code:    127479,
	},
	"sa": Emoji{
		IconURL: fmt.Sprintf("%s/1f202.png?v8", BaseURL),
		Code:    127490,
	},
	"safety_pin": Emoji{
		IconURL: fmt.Sprintf("%s/1f9f7.png?v8", BaseURL),
		Code:    129527,
	},
	"safety_vest": Emoji{
		IconURL: fmt.Sprintf("%s/1f9ba.png?v8", BaseURL),
		Code:    129466,
	},
	"sagittarius": Emoji{
		IconURL: fmt.Sprintf("%s/2650.png?v8", BaseURL),
		Code:    9808,
	},
	"sailboat": Emoji{
		IconURL: fmt.Sprintf("%s/26f5.png?v8", BaseURL),
		Code:    9973,
	},
	"sake": Emoji{
		IconURL: fmt.Sprintf("%s/1f376.png?v8", BaseURL),
		Code:    127862,
	},
	"salt": Emoji{
		IconURL: fmt.Sprintf("%s/1f9c2.png?v8", BaseURL),
		Code:    129474,
	},
	"samoa": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fc-1f1f8.png?v8", BaseURL),
		Code:    127484,
	},
	"san_marino": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1f2.png?v8", BaseURL),
		Code:    127480,
	},
	"sandal": Emoji{
		IconURL: fmt.Sprintf("%s/1f461.png?v8", BaseURL),
		Code:    128097,
	},
	"sandwich": Emoji{
		IconURL: fmt.Sprintf("%s/1f96a.png?v8", BaseURL),
		Code:    129386,
	},
	"santa": Emoji{
		IconURL: fmt.Sprintf("%s/1f385.png?v8", BaseURL),
		Code:    127877,
	},
	"sao_tome_principe": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1f9.png?v8", BaseURL),
		Code:    127480,
	},
	"sari": Emoji{
		IconURL: fmt.Sprintf("%s/1f97b.png?v8", BaseURL),
		Code:    129403,
	},
	"sassy_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f481-2642.png?v8", BaseURL),
		Code:    128129,
	},
	"sassy_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f481-2640.png?v8", BaseURL),
		Code:    128129,
	},
	"satellite": Emoji{
		IconURL: fmt.Sprintf("%s/1f4e1.png?v8", BaseURL),
		Code:    128225,
	},
	"satisfied": Emoji{
		IconURL: fmt.Sprintf("%s/1f606.png?v8", BaseURL),
		Code:    128518,
	},
	"saudi_arabia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1e6.png?v8", BaseURL),
		Code:    127480,
	},
	"sauna_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d6-2642.png?v8", BaseURL),
		Code:    129494,
	},
	"sauna_person": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d6.png?v8", BaseURL),
		Code:    129494,
	},
	"sauna_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d6-2640.png?v8", BaseURL),
		Code:    129494,
	},
	"sauropod": Emoji{
		IconURL: fmt.Sprintf("%s/1f995.png?v8", BaseURL),
		Code:    129429,
	},
	"saxophone": Emoji{
		IconURL: fmt.Sprintf("%s/1f3b7.png?v8", BaseURL),
		Code:    127927,
	},
	"scarf": Emoji{
		IconURL: fmt.Sprintf("%s/1f9e3.png?v8", BaseURL),
		Code:    129507,
	},
	"school": Emoji{
		IconURL: fmt.Sprintf("%s/1f3eb.png?v8", BaseURL),
		Code:    127979,
	},
	"school_satchel": Emoji{
		IconURL: fmt.Sprintf("%s/1f392.png?v8", BaseURL),
		Code:    127890,
	},
	"scientist": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f52c.png?v8", BaseURL),
		Code:    129489,
	},
	"scissors": Emoji{
		IconURL: fmt.Sprintf("%s/2702.png?v8", BaseURL),
		Code:    9986,
	},
	"scorpion": Emoji{
		IconURL: fmt.Sprintf("%s/1f982.png?v8", BaseURL),
		Code:    129410,
	},
	"scorpius": Emoji{
		IconURL: fmt.Sprintf("%s/264f.png?v8", BaseURL),
		Code:    9807,
	},
	"scotland": Emoji{
		IconURL: fmt.Sprintf("%s/1f3f4-e0067-e0062-e0073-e0063-e0074-e007f.png?v8", BaseURL),
		Code:    127988,
	},
	"scream": Emoji{
		IconURL: fmt.Sprintf("%s/1f631.png?v8", BaseURL),
		Code:    128561,
	},
	"scream_cat": Emoji{
		IconURL: fmt.Sprintf("%s/1f640.png?v8", BaseURL),
		Code:    128576,
	},
	"screwdriver": Emoji{
		IconURL: fmt.Sprintf("%s/1fa9b.png?v8", BaseURL),
		Code:    129691,
	},
	"scroll": Emoji{
		IconURL: fmt.Sprintf("%s/1f4dc.png?v8", BaseURL),
		Code:    128220,
	},
	"seal": Emoji{
		IconURL: fmt.Sprintf("%s/1f9ad.png?v8", BaseURL),
		Code:    129453,
	},
	"seat": Emoji{
		IconURL: fmt.Sprintf("%s/1f4ba.png?v8", BaseURL),
		Code:    128186,
	},
	"secret": Emoji{
		IconURL: fmt.Sprintf("%s/3299.png?v8", BaseURL),
		Code:    12953,
	},
	"see_no_evil": Emoji{
		IconURL: fmt.Sprintf("%s/1f648.png?v8", BaseURL),
		Code:    128584,
	},
	"seedling": Emoji{
		IconURL: fmt.Sprintf("%s/1f331.png?v8", BaseURL),
		Code:    127793,
	},
	"selfie": Emoji{
		IconURL: fmt.Sprintf("%s/1f933.png?v8", BaseURL),
		Code:    129331,
	},
	"senegal": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1f3.png?v8", BaseURL),
		Code:    127480,
	},
	"serbia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f7-1f1f8.png?v8", BaseURL),
		Code:    127479,
	},
	"service_dog": Emoji{
		IconURL: fmt.Sprintf("%s/1f415-1f9ba.png?v8", BaseURL),
		Code:    128021,
	},
	"seven": Emoji{
		IconURL: fmt.Sprintf("%s/0037-20e3.png?v8", BaseURL),
		Code:    55,
	},
	"sewing_needle": Emoji{
		IconURL: fmt.Sprintf("%s/1faa1.png?v8", BaseURL),
		Code:    129697,
	},
	"seychelles": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1e8.png?v8", BaseURL),
		Code:    127480,
	},
	"shallow_pan_of_food": Emoji{
		IconURL: fmt.Sprintf("%s/1f958.png?v8", BaseURL),
		Code:    129368,
	},
	"shamrock": Emoji{
		IconURL: fmt.Sprintf("%s/2618.png?v8", BaseURL),
		Code:    9752,
	},
	"shark": Emoji{
		IconURL: fmt.Sprintf("%s/1f988.png?v8", BaseURL),
		Code:    129416,
	},
	"shaved_ice": Emoji{
		IconURL: fmt.Sprintf("%s/1f367.png?v8", BaseURL),
		Code:    127847,
	},
	"sheep": Emoji{
		IconURL: fmt.Sprintf("%s/1f411.png?v8", BaseURL),
		Code:    128017,
	},
	"shell": Emoji{
		IconURL: fmt.Sprintf("%s/1f41a.png?v8", BaseURL),
		Code:    128026,
	},
	"shield": Emoji{
		IconURL: fmt.Sprintf("%s/1f6e1.png?v8", BaseURL),
		Code:    128737,
	},
	"shinto_shrine": Emoji{
		IconURL: fmt.Sprintf("%s/26e9.png?v8", BaseURL),
		Code:    9961,
	},
	"ship": Emoji{
		IconURL: fmt.Sprintf("%s/1f6a2.png?v8", BaseURL),
		Code:    128674,
	},
	"shipit": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/shipit.png?v8",
		Code:    0,
	},
	"shirt": Emoji{
		IconURL: fmt.Sprintf("%s/1f455.png?v8", BaseURL),
		Code:    128085,
	},
	"shit": Emoji{
		IconURL: fmt.Sprintf("%s/1f4a9.png?v8", BaseURL),
		Code:    128169,
	},
	"shoe": Emoji{
		IconURL: fmt.Sprintf("%s/1f45e.png?v8", BaseURL),
		Code:    128094,
	},
	"shopping": Emoji{
		IconURL: fmt.Sprintf("%s/1f6cd.png?v8", BaseURL),
		Code:    128717,
	},
	"shopping_cart": Emoji{
		IconURL: fmt.Sprintf("%s/1f6d2.png?v8", BaseURL),
		Code:    128722,
	},
	"shorts": Emoji{
		IconURL: fmt.Sprintf("%s/1fa73.png?v8", BaseURL),
		Code:    129651,
	},
	"shower": Emoji{
		IconURL: fmt.Sprintf("%s/1f6bf.png?v8", BaseURL),
		Code:    128703,
	},
	"shrimp": Emoji{
		IconURL: fmt.Sprintf("%s/1f990.png?v8", BaseURL),
		Code:    129424,
	},
	"shrug": Emoji{
		IconURL: fmt.Sprintf("%s/1f937.png?v8", BaseURL),
		Code:    129335,
	},
	"shushing_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f92b.png?v8", BaseURL),
		Code:    129323,
	},
	"sierra_leone": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1f1.png?v8", BaseURL),
		Code:    127480,
	},
	"signal_strength": Emoji{
		IconURL: fmt.Sprintf("%s/1f4f6.png?v8", BaseURL),
		Code:    128246,
	},
	"singapore": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1ec.png?v8", BaseURL),
		Code:    127480,
	},
	"singer": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f3a4.png?v8", BaseURL),
		Code:    129489,
	},
	"sint_maarten": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1fd.png?v8", BaseURL),
		Code:    127480,
	},
	"six": Emoji{
		IconURL: fmt.Sprintf("%s/0036-20e3.png?v8", BaseURL),
		Code:    54,
	},
	"six_pointed_star": Emoji{
		IconURL: fmt.Sprintf("%s/1f52f.png?v8", BaseURL),
		Code:    128303,
	},
	"skateboard": Emoji{
		IconURL: fmt.Sprintf("%s/1f6f9.png?v8", BaseURL),
		Code:    128761,
	},
	"ski": Emoji{
		IconURL: fmt.Sprintf("%s/1f3bf.png?v8", BaseURL),
		Code:    127935,
	},
	"skier": Emoji{
		IconURL: fmt.Sprintf("%s/26f7.png?v8", BaseURL),
		Code:    9975,
	},
	"skull": Emoji{
		IconURL: fmt.Sprintf("%s/1f480.png?v8", BaseURL),
		Code:    128128,
	},
	"skull_and_crossbones": Emoji{
		IconURL: fmt.Sprintf("%s/2620.png?v8", BaseURL),
		Code:    9760,
	},
	"skunk": Emoji{
		IconURL: fmt.Sprintf("%s/1f9a8.png?v8", BaseURL),
		Code:    129448,
	},
	"sled": Emoji{
		IconURL: fmt.Sprintf("%s/1f6f7.png?v8", BaseURL),
		Code:    128759,
	},
	"sleeping": Emoji{
		IconURL: fmt.Sprintf("%s/1f634.png?v8", BaseURL),
		Code:    128564,
	},
	"sleeping_bed": Emoji{
		IconURL: fmt.Sprintf("%s/1f6cc.png?v8", BaseURL),
		Code:    128716,
	},
	"sleepy": Emoji{
		IconURL: fmt.Sprintf("%s/1f62a.png?v8", BaseURL),
		Code:    128554,
	},
	"slightly_frowning_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f641.png?v8", BaseURL),
		Code:    128577,
	},
	"slightly_smiling_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f642.png?v8", BaseURL),
		Code:    128578,
	},
	"slot_machine": Emoji{
		IconURL: fmt.Sprintf("%s/1f3b0.png?v8", BaseURL),
		Code:    127920,
	},
	"sloth": Emoji{
		IconURL: fmt.Sprintf("%s/1f9a5.png?v8", BaseURL),
		Code:    129445,
	},
	"slovakia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1f0.png?v8", BaseURL),
		Code:    127480,
	},
	"slovenia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1ee.png?v8", BaseURL),
		Code:    127480,
	},
	"small_airplane": Emoji{
		IconURL: fmt.Sprintf("%s/1f6e9.png?v8", BaseURL),
		Code:    128745,
	},
	"small_blue_diamond": Emoji{
		IconURL: fmt.Sprintf("%s/1f539.png?v8", BaseURL),
		Code:    128313,
	},
	"small_orange_diamond": Emoji{
		IconURL: fmt.Sprintf("%s/1f538.png?v8", BaseURL),
		Code:    128312,
	},
	"small_red_triangle": Emoji{
		IconURL: fmt.Sprintf("%s/1f53a.png?v8", BaseURL),
		Code:    128314,
	},
	"small_red_triangle_down": Emoji{
		IconURL: fmt.Sprintf("%s/1f53b.png?v8", BaseURL),
		Code:    128315,
	},
	"smile": Emoji{
		IconURL: fmt.Sprintf("%s/1f604.png?v8", BaseURL),
		Code:    128516,
	},
	"smile_cat": Emoji{
		IconURL: fmt.Sprintf("%s/1f638.png?v8", BaseURL),
		Code:    128568,
	},
	"smiley": Emoji{
		IconURL: fmt.Sprintf("%s/1f603.png?v8", BaseURL),
		Code:    128515,
	},
	"smiley_cat": Emoji{
		IconURL: fmt.Sprintf("%s/1f63a.png?v8", BaseURL),
		Code:    128570,
	},
	"smiling_face_with_tear": Emoji{
		IconURL: fmt.Sprintf("%s/1f972.png?v8", BaseURL),
		Code:    129394,
	},
	"smiling_face_with_three_hearts": Emoji{
		IconURL: fmt.Sprintf("%s/1f970.png?v8", BaseURL),
		Code:    129392,
	},
	"smiling_imp": Emoji{
		IconURL: fmt.Sprintf("%s/1f608.png?v8", BaseURL),
		Code:    128520,
	},
	"smirk": Emoji{
		IconURL: fmt.Sprintf("%s/1f60f.png?v8", BaseURL),
		Code:    128527,
	},
	"smirk_cat": Emoji{
		IconURL: fmt.Sprintf("%s/1f63c.png?v8", BaseURL),
		Code:    128572,
	},
	"smoking": Emoji{
		IconURL: fmt.Sprintf("%s/1f6ac.png?v8", BaseURL),
		Code:    128684,
	},
	"snail": Emoji{
		IconURL: fmt.Sprintf("%s/1f40c.png?v8", BaseURL),
		Code:    128012,
	},
	"snake": Emoji{
		IconURL: fmt.Sprintf("%s/1f40d.png?v8", BaseURL),
		Code:    128013,
	},
	"sneezing_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f927.png?v8", BaseURL),
		Code:    129319,
	},
	"snowboarder": Emoji{
		IconURL: fmt.Sprintf("%s/1f3c2.png?v8", BaseURL),
		Code:    127938,
	},
	"snowflake": Emoji{
		IconURL: fmt.Sprintf("%s/2744.png?v8", BaseURL),
		Code:    10052,
	},
	"snowman": Emoji{
		IconURL: fmt.Sprintf("%s/26c4.png?v8", BaseURL),
		Code:    9924,
	},
	"snowman_with_snow": Emoji{
		IconURL: fmt.Sprintf("%s/2603.png?v8", BaseURL),
		Code:    9731,
	},
	"soap": Emoji{
		IconURL: fmt.Sprintf("%s/1f9fc.png?v8", BaseURL),
		Code:    129532,
	},
	"sob": Emoji{
		IconURL: fmt.Sprintf("%s/1f62d.png?v8", BaseURL),
		Code:    128557,
	},
	"soccer": Emoji{
		IconURL: fmt.Sprintf("%s/26bd.png?v8", BaseURL),
		Code:    9917,
	},
	"socks": Emoji{
		IconURL: fmt.Sprintf("%s/1f9e6.png?v8", BaseURL),
		Code:    129510,
	},
	"softball": Emoji{
		IconURL: fmt.Sprintf("%s/1f94e.png?v8", BaseURL),
		Code:    129358,
	},
	"solomon_islands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1e7.png?v8", BaseURL),
		Code:    127480,
	},
	"somalia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1f4.png?v8", BaseURL),
		Code:    127480,
	},
	"soon": Emoji{
		IconURL: fmt.Sprintf("%s/1f51c.png?v8", BaseURL),
		Code:    128284,
	},
	"sos": Emoji{
		IconURL: fmt.Sprintf("%s/1f198.png?v8", BaseURL),
		Code:    127384,
	},
	"sound": Emoji{
		IconURL: fmt.Sprintf("%s/1f509.png?v8", BaseURL),
		Code:    128265,
	},
	"south_africa": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ff-1f1e6.png?v8", BaseURL),
		Code:    127487,
	},
	"south_georgia_south_sandwich_islands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1f8.png?v8", BaseURL),
		Code:    127468,
	},
	"south_sudan": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1f8.png?v8", BaseURL),
		Code:    127480,
	},
	"space_invader": Emoji{
		IconURL: fmt.Sprintf("%s/1f47e.png?v8", BaseURL),
		Code:    128126,
	},
	"spades": Emoji{
		IconURL: fmt.Sprintf("%s/2660.png?v8", BaseURL),
		Code:    9824,
	},
	"spaghetti": Emoji{
		IconURL: fmt.Sprintf("%s/1f35d.png?v8", BaseURL),
		Code:    127837,
	},
	"sparkle": Emoji{
		IconURL: fmt.Sprintf("%s/2747.png?v8", BaseURL),
		Code:    10055,
	},
	"sparkler": Emoji{
		IconURL: fmt.Sprintf("%s/1f387.png?v8", BaseURL),
		Code:    127879,
	},
	"sparkles": Emoji{
		IconURL: fmt.Sprintf("%s/2728.png?v8", BaseURL),
		Code:    10024,
	},
	"sparkling_heart": Emoji{
		IconURL: fmt.Sprintf("%s/1f496.png?v8", BaseURL),
		Code:    128150,
	},
	"speak_no_evil": Emoji{
		IconURL: fmt.Sprintf("%s/1f64a.png?v8", BaseURL),
		Code:    128586,
	},
	"speaker": Emoji{
		IconURL: fmt.Sprintf("%s/1f508.png?v8", BaseURL),
		Code:    128264,
	},
	"speaking_head": Emoji{
		IconURL: fmt.Sprintf("%s/1f5e3.png?v8", BaseURL),
		Code:    128483,
	},
	"speech_balloon": Emoji{
		IconURL: fmt.Sprintf("%s/1f4ac.png?v8", BaseURL),
		Code:    128172,
	},
	"speedboat": Emoji{
		IconURL: fmt.Sprintf("%s/1f6a4.png?v8", BaseURL),
		Code:    128676,
	},
	"spider": Emoji{
		IconURL: fmt.Sprintf("%s/1f577.png?v8", BaseURL),
		Code:    128375,
	},
	"spider_web": Emoji{
		IconURL: fmt.Sprintf("%s/1f578.png?v8", BaseURL),
		Code:    128376,
	},
	"spiral_calendar": Emoji{
		IconURL: fmt.Sprintf("%s/1f5d3.png?v8", BaseURL),
		Code:    128467,
	},
	"spiral_notepad": Emoji{
		IconURL: fmt.Sprintf("%s/1f5d2.png?v8", BaseURL),
		Code:    128466,
	},
	"sponge": Emoji{
		IconURL: fmt.Sprintf("%s/1f9fd.png?v8", BaseURL),
		Code:    129533,
	},
	"spoon": Emoji{
		IconURL: fmt.Sprintf("%s/1f944.png?v8", BaseURL),
		Code:    129348,
	},
	"squid": Emoji{
		IconURL: fmt.Sprintf("%s/1f991.png?v8", BaseURL),
		Code:    129425,
	},
	"sri_lanka": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f1-1f1f0.png?v8", BaseURL),
		Code:    127473,
	},
	"st_barthelemy": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e7-1f1f1.png?v8", BaseURL),
		Code:    127463,
	},
	"st_helena": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1ed.png?v8", BaseURL),
		Code:    127480,
	},
	"st_kitts_nevis": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f0-1f1f3.png?v8", BaseURL),
		Code:    127472,
	},
	"st_lucia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f1-1f1e8.png?v8", BaseURL),
		Code:    127473,
	},
	"st_martin": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f2-1f1eb.png?v8", BaseURL),
		Code:    127474,
	},
	"st_pierre_miquelon": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f5-1f1f2.png?v8", BaseURL),
		Code:    127477,
	},
	"st_vincent_grenadines": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fb-1f1e8.png?v8", BaseURL),
		Code:    127483,
	},
	"stadium": Emoji{
		IconURL: fmt.Sprintf("%s/1f3df.png?v8", BaseURL),
		Code:    127967,
	},
	"standing_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f9cd-2642.png?v8", BaseURL),
		Code:    129485,
	},
	"standing_person": Emoji{
		IconURL: fmt.Sprintf("%s/1f9cd.png?v8", BaseURL),
		Code:    129485,
	},
	"standing_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f9cd-2640.png?v8", BaseURL),
		Code:    129485,
	},
	"star": Emoji{
		IconURL: fmt.Sprintf("%s/2b50.png?v8", BaseURL),
		Code:    11088,
	},
	"star2": Emoji{
		IconURL: fmt.Sprintf("%s/1f31f.png?v8", BaseURL),
		Code:    127775,
	},
	"star_and_crescent": Emoji{
		IconURL: fmt.Sprintf("%s/262a.png?v8", BaseURL),
		Code:    9770,
	},
	"star_of_david": Emoji{
		IconURL: fmt.Sprintf("%s/2721.png?v8", BaseURL),
		Code:    10017,
	},
	"star_struck": Emoji{
		IconURL: fmt.Sprintf("%s/1f929.png?v8", BaseURL),
		Code:    129321,
	},
	"stars": Emoji{
		IconURL: fmt.Sprintf("%s/1f320.png?v8", BaseURL),
		Code:    127776,
	},
	"station": Emoji{
		IconURL: fmt.Sprintf("%s/1f689.png?v8", BaseURL),
		Code:    128649,
	},
	"statue_of_liberty": Emoji{
		IconURL: fmt.Sprintf("%s/1f5fd.png?v8", BaseURL),
		Code:    128509,
	},
	"steam_locomotive": Emoji{
		IconURL: fmt.Sprintf("%s/1f682.png?v8", BaseURL),
		Code:    128642,
	},
	"stethoscope": Emoji{
		IconURL: fmt.Sprintf("%s/1fa7a.png?v8", BaseURL),
		Code:    129658,
	},
	"stew": Emoji{
		IconURL: fmt.Sprintf("%s/1f372.png?v8", BaseURL),
		Code:    127858,
	},
	"stop_button": Emoji{
		IconURL: fmt.Sprintf("%s/23f9.png?v8", BaseURL),
		Code:    9209,
	},
	"stop_sign": Emoji{
		IconURL: fmt.Sprintf("%s/1f6d1.png?v8", BaseURL),
		Code:    128721,
	},
	"stopwatch": Emoji{
		IconURL: fmt.Sprintf("%s/23f1.png?v8", BaseURL),
		Code:    9201,
	},
	"straight_ruler": Emoji{
		IconURL: fmt.Sprintf("%s/1f4cf.png?v8", BaseURL),
		Code:    128207,
	},
	"strawberry": Emoji{
		IconURL: fmt.Sprintf("%s/1f353.png?v8", BaseURL),
		Code:    127827,
	},
	"stuck_out_tongue": Emoji{
		IconURL: fmt.Sprintf("%s/1f61b.png?v8", BaseURL),
		Code:    128539,
	},
	"stuck_out_tongue_closed_eyes": Emoji{
		IconURL: fmt.Sprintf("%s/1f61d.png?v8", BaseURL),
		Code:    128541,
	},
	"stuck_out_tongue_winking_eye": Emoji{
		IconURL: fmt.Sprintf("%s/1f61c.png?v8", BaseURL),
		Code:    128540,
	},
	"student": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f393.png?v8", BaseURL),
		Code:    129489,
	},
	"studio_microphone": Emoji{
		IconURL: fmt.Sprintf("%s/1f399.png?v8", BaseURL),
		Code:    127897,
	},
	"stuffed_flatbread": Emoji{
		IconURL: fmt.Sprintf("%s/1f959.png?v8", BaseURL),
		Code:    129369,
	},
	"sudan": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1e9.png?v8", BaseURL),
		Code:    127480,
	},
	"sun_behind_large_cloud": Emoji{
		IconURL: fmt.Sprintf("%s/1f325.png?v8", BaseURL),
		Code:    127781,
	},
	"sun_behind_rain_cloud": Emoji{
		IconURL: fmt.Sprintf("%s/1f326.png?v8", BaseURL),
		Code:    127782,
	},
	"sun_behind_small_cloud": Emoji{
		IconURL: fmt.Sprintf("%s/1f324.png?v8", BaseURL),
		Code:    127780,
	},
	"sun_with_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f31e.png?v8", BaseURL),
		Code:    127774,
	},
	"sunflower": Emoji{
		IconURL: fmt.Sprintf("%s/1f33b.png?v8", BaseURL),
		Code:    127803,
	},
	"sunglasses": Emoji{
		IconURL: fmt.Sprintf("%s/1f60e.png?v8", BaseURL),
		Code:    128526,
	},
	"sunny": Emoji{
		IconURL: fmt.Sprintf("%s/2600.png?v8", BaseURL),
		Code:    9728,
	},
	"sunrise": Emoji{
		IconURL: fmt.Sprintf("%s/1f305.png?v8", BaseURL),
		Code:    127749,
	},
	"sunrise_over_mountains": Emoji{
		IconURL: fmt.Sprintf("%s/1f304.png?v8", BaseURL),
		Code:    127748,
	},
	"superhero": Emoji{
		IconURL: fmt.Sprintf("%s/1f9b8.png?v8", BaseURL),
		Code:    129464,
	},
	"superhero_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f9b8-2642.png?v8", BaseURL),
		Code:    129464,
	},
	"superhero_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f9b8-2640.png?v8", BaseURL),
		Code:    129464,
	},
	"supervillain": Emoji{
		IconURL: fmt.Sprintf("%s/1f9b9.png?v8", BaseURL),
		Code:    129465,
	},
	"supervillain_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f9b9-2642.png?v8", BaseURL),
		Code:    129465,
	},
	"supervillain_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f9b9-2640.png?v8", BaseURL),
		Code:    129465,
	},
	"surfer": Emoji{
		IconURL: fmt.Sprintf("%s/1f3c4.png?v8", BaseURL),
		Code:    127940,
	},
	"surfing_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f3c4-2642.png?v8", BaseURL),
		Code:    127940,
	},
	"surfing_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f3c4-2640.png?v8", BaseURL),
		Code:    127940,
	},
	"suriname": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1f7.png?v8", BaseURL),
		Code:    127480,
	},
	"sushi": Emoji{
		IconURL: fmt.Sprintf("%s/1f363.png?v8", BaseURL),
		Code:    127843,
	},
	"suspect": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/suspect.png?v8",
		Code:    0,
	},
	"suspension_railway": Emoji{
		IconURL: fmt.Sprintf("%s/1f69f.png?v8", BaseURL),
		Code:    128671,
	},
	"svalbard_jan_mayen": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1ef.png?v8", BaseURL),
		Code:    127480,
	},
	"swan": Emoji{
		IconURL: fmt.Sprintf("%s/1f9a2.png?v8", BaseURL),
		Code:    129442,
	},
	"swaziland": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1ff.png?v8", BaseURL),
		Code:    127480,
	},
	"sweat": Emoji{
		IconURL: fmt.Sprintf("%s/1f613.png?v8", BaseURL),
		Code:    128531,
	},
	"sweat_drops": Emoji{
		IconURL: fmt.Sprintf("%s/1f4a6.png?v8", BaseURL),
		Code:    128166,
	},
	"sweat_smile": Emoji{
		IconURL: fmt.Sprintf("%s/1f605.png?v8", BaseURL),
		Code:    128517,
	},
	"sweden": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1ea.png?v8", BaseURL),
		Code:    127480,
	},
	"sweet_potato": Emoji{
		IconURL: fmt.Sprintf("%s/1f360.png?v8", BaseURL),
		Code:    127840,
	},
	"swim_brief": Emoji{
		IconURL: fmt.Sprintf("%s/1fa72.png?v8", BaseURL),
		Code:    129650,
	},
	"swimmer": Emoji{
		IconURL: fmt.Sprintf("%s/1f3ca.png?v8", BaseURL),
		Code:    127946,
	},
	"swimming_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f3ca-2642.png?v8", BaseURL),
		Code:    127946,
	},
	"swimming_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f3ca-2640.png?v8", BaseURL),
		Code:    127946,
	},
	"switzerland": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e8-1f1ed.png?v8", BaseURL),
		Code:    127464,
	},
	"symbols": Emoji{
		IconURL: fmt.Sprintf("%s/1f523.png?v8", BaseURL),
		Code:    128291,
	},
	"synagogue": Emoji{
		IconURL: fmt.Sprintf("%s/1f54d.png?v8", BaseURL),
		Code:    128333,
	},
	"syria": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f8-1f1fe.png?v8", BaseURL),
		Code:    127480,
	},
	"syringe": Emoji{
		IconURL: fmt.Sprintf("%s/1f489.png?v8", BaseURL),
		Code:    128137,
	},
	"t-rex": Emoji{
		IconURL: fmt.Sprintf("%s/1f996.png?v8", BaseURL),
		Code:    129430,
	},
	"taco": Emoji{
		IconURL: fmt.Sprintf("%s/1f32e.png?v8", BaseURL),
		Code:    127790,
	},
	"tada": Emoji{
		IconURL: fmt.Sprintf("%s/1f389.png?v8", BaseURL),
		Code:    127881,
	},
	"taiwan": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f9-1f1fc.png?v8", BaseURL),
		Code:    127481,
	},
	"tajikistan": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f9-1f1ef.png?v8", BaseURL),
		Code:    127481,
	},
	"takeout_box": Emoji{
		IconURL: fmt.Sprintf("%s/1f961.png?v8", BaseURL),
		Code:    129377,
	},
	"tamale": Emoji{
		IconURL: fmt.Sprintf("%s/1fad4.png?v8", BaseURL),
		Code:    129748,
	},
	"tanabata_tree": Emoji{
		IconURL: fmt.Sprintf("%s/1f38b.png?v8", BaseURL),
		Code:    127883,
	},
	"tangerine": Emoji{
		IconURL: fmt.Sprintf("%s/1f34a.png?v8", BaseURL),
		Code:    127818,
	},
	"tanzania": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f9-1f1ff.png?v8", BaseURL),
		Code:    127481,
	},
	"taurus": Emoji{
		IconURL: fmt.Sprintf("%s/2649.png?v8", BaseURL),
		Code:    9801,
	},
	"taxi": Emoji{
		IconURL: fmt.Sprintf("%s/1f695.png?v8", BaseURL),
		Code:    128661,
	},
	"tea": Emoji{
		IconURL: fmt.Sprintf("%s/1f375.png?v8", BaseURL),
		Code:    127861,
	},
	"teacher": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f3eb.png?v8", BaseURL),
		Code:    129489,
	},
	"teapot": Emoji{
		IconURL: fmt.Sprintf("%s/1fad6.png?v8", BaseURL),
		Code:    129750,
	},
	"technologist": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d1-1f4bb.png?v8", BaseURL),
		Code:    129489,
	},
	"teddy_bear": Emoji{
		IconURL: fmt.Sprintf("%s/1f9f8.png?v8", BaseURL),
		Code:    129528,
	},
	"telephone": Emoji{
		IconURL: fmt.Sprintf("%s/260e.png?v8", BaseURL),
		Code:    9742,
	},
	"telephone_receiver": Emoji{
		IconURL: fmt.Sprintf("%s/1f4de.png?v8", BaseURL),
		Code:    128222,
	},
	"telescope": Emoji{
		IconURL: fmt.Sprintf("%s/1f52d.png?v8", BaseURL),
		Code:    128301,
	},
	"tennis": Emoji{
		IconURL: fmt.Sprintf("%s/1f3be.png?v8", BaseURL),
		Code:    127934,
	},
	"tent": Emoji{
		IconURL: fmt.Sprintf("%s/26fa.png?v8", BaseURL),
		Code:    9978,
	},
	"test_tube": Emoji{
		IconURL: fmt.Sprintf("%s/1f9ea.png?v8", BaseURL),
		Code:    129514,
	},
	"thailand": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f9-1f1ed.png?v8", BaseURL),
		Code:    127481,
	},
	"thermometer": Emoji{
		IconURL: fmt.Sprintf("%s/1f321.png?v8", BaseURL),
		Code:    127777,
	},
	"thinking": Emoji{
		IconURL: fmt.Sprintf("%s/1f914.png?v8", BaseURL),
		Code:    129300,
	},
	"thong_sandal": Emoji{
		IconURL: fmt.Sprintf("%s/1fa74.png?v8", BaseURL),
		Code:    129652,
	},
	"thought_balloon": Emoji{
		IconURL: fmt.Sprintf("%s/1f4ad.png?v8", BaseURL),
		Code:    128173,
	},
	"thread": Emoji{
		IconURL: fmt.Sprintf("%s/1f9f5.png?v8", BaseURL),
		Code:    129525,
	},
	"three": Emoji{
		IconURL: fmt.Sprintf("%s/0033-20e3.png?v8", BaseURL),
		Code:    51,
	},
	"thumbsdown": Emoji{
		IconURL: fmt.Sprintf("%s/1f44e.png?v8", BaseURL),
		Code:    128078,
	},
	"thumbsup": Emoji{
		IconURL: fmt.Sprintf("%s/1f44d.png?v8", BaseURL),
		Code:    128077,
	},
	"ticket": Emoji{
		IconURL: fmt.Sprintf("%s/1f3ab.png?v8", BaseURL),
		Code:    127915,
	},
	"tickets": Emoji{
		IconURL: fmt.Sprintf("%s/1f39f.png?v8", BaseURL),
		Code:    127903,
	},
	"tiger": Emoji{
		IconURL: fmt.Sprintf("%s/1f42f.png?v8", BaseURL),
		Code:    128047,
	},
	"tiger2": Emoji{
		IconURL: fmt.Sprintf("%s/1f405.png?v8", BaseURL),
		Code:    128005,
	},
	"timer_clock": Emoji{
		IconURL: fmt.Sprintf("%s/23f2.png?v8", BaseURL),
		Code:    9202,
	},
	"timor_leste": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f9-1f1f1.png?v8", BaseURL),
		Code:    127481,
	},
	"tipping_hand_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f481-2642.png?v8", BaseURL),
		Code:    128129,
	},
	"tipping_hand_person": Emoji{
		IconURL: fmt.Sprintf("%s/1f481.png?v8", BaseURL),
		Code:    128129,
	},
	"tipping_hand_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f481-2640.png?v8", BaseURL),
		Code:    128129,
	},
	"tired_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f62b.png?v8", BaseURL),
		Code:    128555,
	},
	"tm": Emoji{
		IconURL: fmt.Sprintf("%s/2122.png?v8", BaseURL),
		Code:    8482,
	},
	"togo": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f9-1f1ec.png?v8", BaseURL),
		Code:    127481,
	},
	"toilet": Emoji{
		IconURL: fmt.Sprintf("%s/1f6bd.png?v8", BaseURL),
		Code:    128701,
	},
	"tokelau": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f9-1f1f0.png?v8", BaseURL),
		Code:    127481,
	},
	"tokyo_tower": Emoji{
		IconURL: fmt.Sprintf("%s/1f5fc.png?v8", BaseURL),
		Code:    128508,
	},
	"tomato": Emoji{
		IconURL: fmt.Sprintf("%s/1f345.png?v8", BaseURL),
		Code:    127813,
	},
	"tonga": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f9-1f1f4.png?v8", BaseURL),
		Code:    127481,
	},
	"tongue": Emoji{
		IconURL: fmt.Sprintf("%s/1f445.png?v8", BaseURL),
		Code:    128069,
	},
	"toolbox": Emoji{
		IconURL: fmt.Sprintf("%s/1f9f0.png?v8", BaseURL),
		Code:    129520,
	},
	"tooth": Emoji{
		IconURL: fmt.Sprintf("%s/1f9b7.png?v8", BaseURL),
		Code:    129463,
	},
	"toothbrush": Emoji{
		IconURL: fmt.Sprintf("%s/1faa5.png?v8", BaseURL),
		Code:    129701,
	},
	"top": Emoji{
		IconURL: fmt.Sprintf("%s/1f51d.png?v8", BaseURL),
		Code:    128285,
	},
	"tophat": Emoji{
		IconURL: fmt.Sprintf("%s/1f3a9.png?v8", BaseURL),
		Code:    127913,
	},
	"tornado": Emoji{
		IconURL: fmt.Sprintf("%s/1f32a.png?v8", BaseURL),
		Code:    127786,
	},
	"tr": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f9-1f1f7.png?v8", BaseURL),
		Code:    127481,
	},
	"trackball": Emoji{
		IconURL: fmt.Sprintf("%s/1f5b2.png?v8", BaseURL),
		Code:    128434,
	},
	"tractor": Emoji{
		IconURL: fmt.Sprintf("%s/1f69c.png?v8", BaseURL),
		Code:    128668,
	},
	"traffic_light": Emoji{
		IconURL: fmt.Sprintf("%s/1f6a5.png?v8", BaseURL),
		Code:    128677,
	},
	"train": Emoji{
		IconURL: fmt.Sprintf("%s/1f68b.png?v8", BaseURL),
		Code:    128651,
	},
	"train2": Emoji{
		IconURL: fmt.Sprintf("%s/1f686.png?v8", BaseURL),
		Code:    128646,
	},
	"tram": Emoji{
		IconURL: fmt.Sprintf("%s/1f68a.png?v8", BaseURL),
		Code:    128650,
	},
	"transgender_flag": Emoji{
		IconURL: fmt.Sprintf("%s/1f3f3-26a7.png?v8", BaseURL),
		Code:    127987,
	},
	"transgender_symbol": Emoji{
		IconURL: fmt.Sprintf("%s/26a7.png?v8", BaseURL),
		Code:    9895,
	},
	"triangular_flag_on_post": Emoji{
		IconURL: fmt.Sprintf("%s/1f6a9.png?v8", BaseURL),
		Code:    128681,
	},
	"triangular_ruler": Emoji{
		IconURL: fmt.Sprintf("%s/1f4d0.png?v8", BaseURL),
		Code:    128208,
	},
	"trident": Emoji{
		IconURL: fmt.Sprintf("%s/1f531.png?v8", BaseURL),
		Code:    128305,
	},
	"trinidad_tobago": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f9-1f1f9.png?v8", BaseURL),
		Code:    127481,
	},
	"tristan_da_cunha": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f9-1f1e6.png?v8", BaseURL),
		Code:    127481,
	},
	"triumph": Emoji{
		IconURL: fmt.Sprintf("%s/1f624.png?v8", BaseURL),
		Code:    128548,
	},
	"trolleybus": Emoji{
		IconURL: fmt.Sprintf("%s/1f68e.png?v8", BaseURL),
		Code:    128654,
	},
	"trollface": Emoji{
		IconURL: "https://github.githubassets.com/images/icons/emoji/trollface.png?v8",
		Code:    0,
	},
	"trophy": Emoji{
		IconURL: fmt.Sprintf("%s/1f3c6.png?v8", BaseURL),
		Code:    127942,
	},
	"tropical_drink": Emoji{
		IconURL: fmt.Sprintf("%s/1f379.png?v8", BaseURL),
		Code:    127865,
	},
	"tropical_fish": Emoji{
		IconURL: fmt.Sprintf("%s/1f420.png?v8", BaseURL),
		Code:    128032,
	},
	"truck": Emoji{
		IconURL: fmt.Sprintf("%s/1f69a.png?v8", BaseURL),
		Code:    128666,
	},
	"trumpet": Emoji{
		IconURL: fmt.Sprintf("%s/1f3ba.png?v8", BaseURL),
		Code:    127930,
	},
	"tshirt": Emoji{
		IconURL: fmt.Sprintf("%s/1f455.png?v8", BaseURL),
		Code:    128085,
	},
	"tulip": Emoji{
		IconURL: fmt.Sprintf("%s/1f337.png?v8", BaseURL),
		Code:    127799,
	},
	"tumbler_glass": Emoji{
		IconURL: fmt.Sprintf("%s/1f943.png?v8", BaseURL),
		Code:    129347,
	},
	"tunisia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f9-1f1f3.png?v8", BaseURL),
		Code:    127481,
	},
	"turkey": Emoji{
		IconURL: fmt.Sprintf("%s/1f983.png?v8", BaseURL),
		Code:    129411,
	},
	"turkmenistan": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f9-1f1f2.png?v8", BaseURL),
		Code:    127481,
	},
	"turks_caicos_islands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f9-1f1e8.png?v8", BaseURL),
		Code:    127481,
	},
	"turtle": Emoji{
		IconURL: fmt.Sprintf("%s/1f422.png?v8", BaseURL),
		Code:    128034,
	},
	"tuvalu": Emoji{
		IconURL: fmt.Sprintf("%s/1f1f9-1f1fb.png?v8", BaseURL),
		Code:    127481,
	},
	"tv": Emoji{
		IconURL: fmt.Sprintf("%s/1f4fa.png?v8", BaseURL),
		Code:    128250,
	},
	"twisted_rightwards_arrows": Emoji{
		IconURL: fmt.Sprintf("%s/1f500.png?v8", BaseURL),
		Code:    128256,
	},
	"two": Emoji{
		IconURL: fmt.Sprintf("%s/0032-20e3.png?v8", BaseURL),
		Code:    50,
	},
	"two_hearts": Emoji{
		IconURL: fmt.Sprintf("%s/1f495.png?v8", BaseURL),
		Code:    128149,
	},
	"two_men_holding_hands": Emoji{
		IconURL: fmt.Sprintf("%s/1f46c.png?v8", BaseURL),
		Code:    128108,
	},
	"two_women_holding_hands": Emoji{
		IconURL: fmt.Sprintf("%s/1f46d.png?v8", BaseURL),
		Code:    128109,
	},
	"u5272": Emoji{
		IconURL: fmt.Sprintf("%s/1f239.png?v8", BaseURL),
		Code:    127545,
	},
	"u5408": Emoji{
		IconURL: fmt.Sprintf("%s/1f234.png?v8", BaseURL),
		Code:    127540,
	},
	"u55b6": Emoji{
		IconURL: fmt.Sprintf("%s/1f23a.png?v8", BaseURL),
		Code:    127546,
	},
	"u6307": Emoji{
		IconURL: fmt.Sprintf("%s/1f22f.png?v8", BaseURL),
		Code:    127535,
	},
	"u6708": Emoji{
		IconURL: fmt.Sprintf("%s/1f237.png?v8", BaseURL),
		Code:    127543,
	},
	"u6709": Emoji{
		IconURL: fmt.Sprintf("%s/1f236.png?v8", BaseURL),
		Code:    127542,
	},
	"u6e80": Emoji{
		IconURL: fmt.Sprintf("%s/1f235.png?v8", BaseURL),
		Code:    127541,
	},
	"u7121": Emoji{
		IconURL: fmt.Sprintf("%s/1f21a.png?v8", BaseURL),
		Code:    127514,
	},
	"u7533": Emoji{
		IconURL: fmt.Sprintf("%s/1f238.png?v8", BaseURL),
		Code:    127544,
	},
	"u7981": Emoji{
		IconURL: fmt.Sprintf("%s/1f232.png?v8", BaseURL),
		Code:    127538,
	},
	"u7a7a": Emoji{
		IconURL: fmt.Sprintf("%s/1f233.png?v8", BaseURL),
		Code:    127539,
	},
	"uganda": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fa-1f1ec.png?v8", BaseURL),
		Code:    127482,
	},
	"uk": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ec-1f1e7.png?v8", BaseURL),
		Code:    127468,
	},
	"ukraine": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fa-1f1e6.png?v8", BaseURL),
		Code:    127482,
	},
	"umbrella": Emoji{
		IconURL: fmt.Sprintf("%s/2614.png?v8", BaseURL),
		Code:    9748,
	},
	"unamused": Emoji{
		IconURL: fmt.Sprintf("%s/1f612.png?v8", BaseURL),
		Code:    128530,
	},
	"underage": Emoji{
		IconURL: fmt.Sprintf("%s/1f51e.png?v8", BaseURL),
		Code:    128286,
	},
	"unicorn": Emoji{
		IconURL: fmt.Sprintf("%s/1f984.png?v8", BaseURL),
		Code:    129412,
	},
	"united_arab_emirates": Emoji{
		IconURL: fmt.Sprintf("%s/1f1e6-1f1ea.png?v8", BaseURL),
		Code:    127462,
	},
	"united_nations": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fa-1f1f3.png?v8", BaseURL),
		Code:    127482,
	},
	"unlock": Emoji{
		IconURL: fmt.Sprintf("%s/1f513.png?v8", BaseURL),
		Code:    128275,
	},
	"up": Emoji{
		IconURL: fmt.Sprintf("%s/1f199.png?v8", BaseURL),
		Code:    127385,
	},
	"upside_down_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f643.png?v8", BaseURL),
		Code:    128579,
	},
	"uruguay": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fa-1f1fe.png?v8", BaseURL),
		Code:    127482,
	},
	"us": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fa-1f1f8.png?v8", BaseURL),
		Code:    127482,
	},
	"us_outlying_islands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fa-1f1f2.png?v8", BaseURL),
		Code:    127482,
	},
	"us_virgin_islands": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fb-1f1ee.png?v8", BaseURL),
		Code:    127483,
	},
	"uzbekistan": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fa-1f1ff.png?v8", BaseURL),
		Code:    127482,
	},
	"v": Emoji{
		IconURL: fmt.Sprintf("%s/270c.png?v8", BaseURL),
		Code:    9996,
	},
	"vampire": Emoji{
		IconURL: fmt.Sprintf("%s/1f9db.png?v8", BaseURL),
		Code:    129499,
	},
	"vampire_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f9db-2642.png?v8", BaseURL),
		Code:    129499,
	},
	"vampire_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f9db-2640.png?v8", BaseURL),
		Code:    129499,
	},
	"vanuatu": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fb-1f1fa.png?v8", BaseURL),
		Code:    127483,
	},
	"vatican_city": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fb-1f1e6.png?v8", BaseURL),
		Code:    127483,
	},
	"venezuela": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fb-1f1ea.png?v8", BaseURL),
		Code:    127483,
	},
	"vertical_traffic_light": Emoji{
		IconURL: fmt.Sprintf("%s/1f6a6.png?v8", BaseURL),
		Code:    128678,
	},
	"vhs": Emoji{
		IconURL: fmt.Sprintf("%s/1f4fc.png?v8", BaseURL),
		Code:    128252,
	},
	"vibration_mode": Emoji{
		IconURL: fmt.Sprintf("%s/1f4f3.png?v8", BaseURL),
		Code:    128243,
	},
	"video_camera": Emoji{
		IconURL: fmt.Sprintf("%s/1f4f9.png?v8", BaseURL),
		Code:    128249,
	},
	"video_game": Emoji{
		IconURL: fmt.Sprintf("%s/1f3ae.png?v8", BaseURL),
		Code:    127918,
	},
	"vietnam": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fb-1f1f3.png?v8", BaseURL),
		Code:    127483,
	},
	"violin": Emoji{
		IconURL: fmt.Sprintf("%s/1f3bb.png?v8", BaseURL),
		Code:    127931,
	},
	"virgo": Emoji{
		IconURL: fmt.Sprintf("%s/264d.png?v8", BaseURL),
		Code:    9805,
	},
	"volcano": Emoji{
		IconURL: fmt.Sprintf("%s/1f30b.png?v8", BaseURL),
		Code:    127755,
	},
	"volleyball": Emoji{
		IconURL: fmt.Sprintf("%s/1f3d0.png?v8", BaseURL),
		Code:    127952,
	},
	"vomiting_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f92e.png?v8", BaseURL),
		Code:    129326,
	},
	"vs": Emoji{
		IconURL: fmt.Sprintf("%s/1f19a.png?v8", BaseURL),
		Code:    127386,
	},
	"vulcan_salute": Emoji{
		IconURL: fmt.Sprintf("%s/1f596.png?v8", BaseURL),
		Code:    128406,
	},
	"waffle": Emoji{
		IconURL: fmt.Sprintf("%s/1f9c7.png?v8", BaseURL),
		Code:    129479,
	},
	"wales": Emoji{
		IconURL: fmt.Sprintf("%s/1f3f4-e0067-e0062-e0077-e006c-e0073-e007f.png?v8", BaseURL),
		Code:    127988,
	},
	"walking": Emoji{
		IconURL: fmt.Sprintf("%s/1f6b6.png?v8", BaseURL),
		Code:    128694,
	},
	"walking_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f6b6-2642.png?v8", BaseURL),
		Code:    128694,
	},
	"walking_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f6b6-2640.png?v8", BaseURL),
		Code:    128694,
	},
	"wallis_futuna": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fc-1f1eb.png?v8", BaseURL),
		Code:    127484,
	},
	"waning_crescent_moon": Emoji{
		IconURL: fmt.Sprintf("%s/1f318.png?v8", BaseURL),
		Code:    127768,
	},
	"waning_gibbous_moon": Emoji{
		IconURL: fmt.Sprintf("%s/1f316.png?v8", BaseURL),
		Code:    127766,
	},
	"warning": Emoji{
		IconURL: fmt.Sprintf("%s/26a0.png?v8", BaseURL),
		Code:    9888,
	},
	"wastebasket": Emoji{
		IconURL: fmt.Sprintf("%s/1f5d1.png?v8", BaseURL),
		Code:    128465,
	},
	"watch": Emoji{
		IconURL: fmt.Sprintf("%s/231a.png?v8", BaseURL),
		Code:    8986,
	},
	"water_buffalo": Emoji{
		IconURL: fmt.Sprintf("%s/1f403.png?v8", BaseURL),
		Code:    128003,
	},
	"water_polo": Emoji{
		IconURL: fmt.Sprintf("%s/1f93d.png?v8", BaseURL),
		Code:    129341,
	},
	"watermelon": Emoji{
		IconURL: fmt.Sprintf("%s/1f349.png?v8", BaseURL),
		Code:    127817,
	},
	"wave": Emoji{
		IconURL: fmt.Sprintf("%s/1f44b.png?v8", BaseURL),
		Code:    128075,
	},
	"wavy_dash": Emoji{
		IconURL: fmt.Sprintf("%s/3030.png?v8", BaseURL),
		Code:    12336,
	},
	"waxing_crescent_moon": Emoji{
		IconURL: fmt.Sprintf("%s/1f312.png?v8", BaseURL),
		Code:    127762,
	},
	"waxing_gibbous_moon": Emoji{
		IconURL: fmt.Sprintf("%s/1f314.png?v8", BaseURL),
		Code:    127764,
	},
	"wc": Emoji{
		IconURL: fmt.Sprintf("%s/1f6be.png?v8", BaseURL),
		Code:    128702,
	},
	"weary": Emoji{
		IconURL: fmt.Sprintf("%s/1f629.png?v8", BaseURL),
		Code:    128553,
	},
	"wedding": Emoji{
		IconURL: fmt.Sprintf("%s/1f492.png?v8", BaseURL),
		Code:    128146,
	},
	"weight_lifting": Emoji{
		IconURL: fmt.Sprintf("%s/1f3cb.png?v8", BaseURL),
		Code:    127947,
	},
	"weight_lifting_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f3cb-2642.png?v8", BaseURL),
		Code:    127947,
	},
	"weight_lifting_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f3cb-2640.png?v8", BaseURL),
		Code:    127947,
	},
	"western_sahara": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ea-1f1ed.png?v8", BaseURL),
		Code:    127466,
	},
	"whale": Emoji{
		IconURL: fmt.Sprintf("%s/1f433.png?v8", BaseURL),
		Code:    128051,
	},
	"whale2": Emoji{
		IconURL: fmt.Sprintf("%s/1f40b.png?v8", BaseURL),
		Code:    128011,
	},
	"wheel_of_dharma": Emoji{
		IconURL: fmt.Sprintf("%s/2638.png?v8", BaseURL),
		Code:    9784,
	},
	"wheelchair": Emoji{
		IconURL: fmt.Sprintf("%s/267f.png?v8", BaseURL),
		Code:    9855,
	},
	"white_check_mark": Emoji{
		IconURL: fmt.Sprintf("%s/2705.png?v8", BaseURL),
		Code:    9989,
	},
	"white_circle": Emoji{
		IconURL: fmt.Sprintf("%s/26aa.png?v8", BaseURL),
		Code:    9898,
	},
	"white_flag": Emoji{
		IconURL: fmt.Sprintf("%s/1f3f3.png?v8", BaseURL),
		Code:    127987,
	},
	"white_flower": Emoji{
		IconURL: fmt.Sprintf("%s/1f4ae.png?v8", BaseURL),
		Code:    128174,
	},
	"white_haired_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f468-1f9b3.png?v8", BaseURL),
		Code:    128104,
	},
	"white_haired_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f9b3.png?v8", BaseURL),
		Code:    128105,
	},
	"white_heart": Emoji{
		IconURL: fmt.Sprintf("%s/1f90d.png?v8", BaseURL),
		Code:    129293,
	},
	"white_large_square": Emoji{
		IconURL: fmt.Sprintf("%s/2b1c.png?v8", BaseURL),
		Code:    11036,
	},
	"white_medium_small_square": Emoji{
		IconURL: fmt.Sprintf("%s/25fd.png?v8", BaseURL),
		Code:    9725,
	},
	"white_medium_square": Emoji{
		IconURL: fmt.Sprintf("%s/25fb.png?v8", BaseURL),
		Code:    9723,
	},
	"white_small_square": Emoji{
		IconURL: fmt.Sprintf("%s/25ab.png?v8", BaseURL),
		Code:    9643,
	},
	"white_square_button": Emoji{
		IconURL: fmt.Sprintf("%s/1f533.png?v8", BaseURL),
		Code:    128307,
	},
	"wilted_flower": Emoji{
		IconURL: fmt.Sprintf("%s/1f940.png?v8", BaseURL),
		Code:    129344,
	},
	"wind_chime": Emoji{
		IconURL: fmt.Sprintf("%s/1f390.png?v8", BaseURL),
		Code:    127888,
	},
	"wind_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f32c.png?v8", BaseURL),
		Code:    127788,
	},
	"window": Emoji{
		IconURL: fmt.Sprintf("%s/1fa9f.png?v8", BaseURL),
		Code:    129695,
	},
	"wine_glass": Emoji{
		IconURL: fmt.Sprintf("%s/1f377.png?v8", BaseURL),
		Code:    127863,
	},
	"wink": Emoji{
		IconURL: fmt.Sprintf("%s/1f609.png?v8", BaseURL),
		Code:    128521,
	},
	"wolf": Emoji{
		IconURL: fmt.Sprintf("%s/1f43a.png?v8", BaseURL),
		Code:    128058,
	},
	"woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f469.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_artist": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f3a8.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_astronaut": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f680.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_beard": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d4-2640.png?v8", BaseURL),
		Code:    129492,
	},
	"woman_cartwheeling": Emoji{
		IconURL: fmt.Sprintf("%s/1f938-2640.png?v8", BaseURL),
		Code:    129336,
	},
	"woman_cook": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f373.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_dancing": Emoji{
		IconURL: fmt.Sprintf("%s/1f483.png?v8", BaseURL),
		Code:    128131,
	},
	"woman_facepalming": Emoji{
		IconURL: fmt.Sprintf("%s/1f926-2640.png?v8", BaseURL),
		Code:    129318,
	},
	"woman_factory_worker": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f3ed.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_farmer": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f33e.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_feeding_baby": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f37c.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_firefighter": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f692.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_health_worker": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-2695.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_in_manual_wheelchair": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f9bd.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_in_motorized_wheelchair": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f9bc.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_in_tuxedo": Emoji{
		IconURL: fmt.Sprintf("%s/1f935-2640.png?v8", BaseURL),
		Code:    129333,
	},
	"woman_judge": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-2696.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_juggling": Emoji{
		IconURL: fmt.Sprintf("%s/1f939-2640.png?v8", BaseURL),
		Code:    129337,
	},
	"woman_mechanic": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f527.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_office_worker": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f4bc.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_pilot": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-2708.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_playing_handball": Emoji{
		IconURL: fmt.Sprintf("%s/1f93e-2640.png?v8", BaseURL),
		Code:    129342,
	},
	"woman_playing_water_polo": Emoji{
		IconURL: fmt.Sprintf("%s/1f93d-2640.png?v8", BaseURL),
		Code:    129341,
	},
	"woman_scientist": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f52c.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_shrugging": Emoji{
		IconURL: fmt.Sprintf("%s/1f937-2640.png?v8", BaseURL),
		Code:    129335,
	},
	"woman_singer": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f3a4.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_student": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f393.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_teacher": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f3eb.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_technologist": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f4bb.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_with_headscarf": Emoji{
		IconURL: fmt.Sprintf("%s/1f9d5.png?v8", BaseURL),
		Code:    129493,
	},
	"woman_with_probing_cane": Emoji{
		IconURL: fmt.Sprintf("%s/1f469-1f9af.png?v8", BaseURL),
		Code:    128105,
	},
	"woman_with_turban": Emoji{
		IconURL: fmt.Sprintf("%s/1f473-2640.png?v8", BaseURL),
		Code:    128115,
	},
	"woman_with_veil": Emoji{
		IconURL: fmt.Sprintf("%s/1f470-2640.png?v8", BaseURL),
		Code:    128112,
	},
	"womans_clothes": Emoji{
		IconURL: fmt.Sprintf("%s/1f45a.png?v8", BaseURL),
		Code:    128090,
	},
	"womans_hat": Emoji{
		IconURL: fmt.Sprintf("%s/1f452.png?v8", BaseURL),
		Code:    128082,
	},
	"women_wrestling": Emoji{
		IconURL: fmt.Sprintf("%s/1f93c-2640.png?v8", BaseURL),
		Code:    129340,
	},
	"womens": Emoji{
		IconURL: fmt.Sprintf("%s/1f6ba.png?v8", BaseURL),
		Code:    128698,
	},
	"wood": Emoji{
		IconURL: fmt.Sprintf("%s/1fab5.png?v8", BaseURL),
		Code:    129717,
	},
	"woozy_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f974.png?v8", BaseURL),
		Code:    129396,
	},
	"world_map": Emoji{
		IconURL: fmt.Sprintf("%s/1f5fa.png?v8", BaseURL),
		Code:    128506,
	},
	"worm": Emoji{
		IconURL: fmt.Sprintf("%s/1fab1.png?v8", BaseURL),
		Code:    129713,
	},
	"worried": Emoji{
		IconURL: fmt.Sprintf("%s/1f61f.png?v8", BaseURL),
		Code:    128543,
	},
	"wrench": Emoji{
		IconURL: fmt.Sprintf("%s/1f527.png?v8", BaseURL),
		Code:    128295,
	},
	"wrestling": Emoji{
		IconURL: fmt.Sprintf("%s/1f93c.png?v8", BaseURL),
		Code:    129340,
	},
	"writing_hand": Emoji{
		IconURL: fmt.Sprintf("%s/270d.png?v8", BaseURL),
		Code:    9997,
	},
	"x": Emoji{
		IconURL: fmt.Sprintf("%s/274c.png?v8", BaseURL),
		Code:    10060,
	},
	"yarn": Emoji{
		IconURL: fmt.Sprintf("%s/1f9f6.png?v8", BaseURL),
		Code:    129526,
	},
	"yawning_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f971.png?v8", BaseURL),
		Code:    129393,
	},
	"yellow_circle": Emoji{
		IconURL: fmt.Sprintf("%s/1f7e1.png?v8", BaseURL),
		Code:    128993,
	},
	"yellow_heart": Emoji{
		IconURL: fmt.Sprintf("%s/1f49b.png?v8", BaseURL),
		Code:    128155,
	},
	"yellow_square": Emoji{
		IconURL: fmt.Sprintf("%s/1f7e8.png?v8", BaseURL),
		Code:    129000,
	},
	"yemen": Emoji{
		IconURL: fmt.Sprintf("%s/1f1fe-1f1ea.png?v8", BaseURL),
		Code:    127486,
	},
	"yen": Emoji{
		IconURL: fmt.Sprintf("%s/1f4b4.png?v8", BaseURL),
		Code:    128180,
	},
	"yin_yang": Emoji{
		IconURL: fmt.Sprintf("%s/262f.png?v8", BaseURL),
		Code:    9775,
	},
	"yo_yo": Emoji{
		IconURL: fmt.Sprintf("%s/1fa80.png?v8", BaseURL),
		Code:    129664,
	},
	"yum": Emoji{
		IconURL: fmt.Sprintf("%s/1f60b.png?v8", BaseURL),
		Code:    128523,
	},
	"zambia": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ff-1f1f2.png?v8", BaseURL),
		Code:    127487,
	},
	"zany_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f92a.png?v8", BaseURL),
		Code:    129322,
	},
	"zap": Emoji{
		IconURL: fmt.Sprintf("%s/26a1.png?v8", BaseURL),
		Code:    9889,
	},
	"zebra": Emoji{
		IconURL: fmt.Sprintf("%s/1f993.png?v8", BaseURL),
		Code:    129427,
	},
	"zero": Emoji{
		IconURL: fmt.Sprintf("%s/0030-20e3.png?v8", BaseURL),
		Code:    48,
	},
	"zimbabwe": Emoji{
		IconURL: fmt.Sprintf("%s/1f1ff-1f1fc.png?v8", BaseURL),
		Code:    127487,
	},
	"zipper_mouth_face": Emoji{
		IconURL: fmt.Sprintf("%s/1f910.png?v8", BaseURL),
		Code:    129296,
	},
	"zombie": Emoji{
		IconURL: fmt.Sprintf("%s/1f9df.png?v8", BaseURL),
		Code:    129503,
	},
	"zombie_man": Emoji{
		IconURL: fmt.Sprintf("%s/1f9df-2642.png?v8", BaseURL),
		Code:    129503,
	},
	"zombie_woman": Emoji{
		IconURL: fmt.Sprintf("%s/1f9df-2640.png?v8", BaseURL),
		Code:    129503,
	},
	"zzz": Emoji{
		IconURL: fmt.Sprintf("%s/1f4a4.png?v8", BaseURL),
		Code:    128164,
	},
}

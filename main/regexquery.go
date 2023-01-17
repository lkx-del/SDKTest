package main

import (
	"fmt"
	"github.com/influxdata/influxdb1-client/v2"
	"time"
)

func main() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: "admin",
		Password: "At1314comi!",
	})
	if err != nil {
		fmt.Println("Error", err.Error())
	}
	defer c.Close()

	/*mgram := make([]string, 30)
	mgram[0] = "select * from clvTable where logs match 'GET'"
	mgram[1] = "select * from clvTable where logs match 'tickets/'"
	mgram[2] = "select * from clvTable where logs match 'playing'"
	mgram[3] = "select * from clvTable where logs match 'GET /images/'"
	mgram[4] = "select * from clvTable where logs match 'cal_mont.gif'"
	mgram[5] = "select * from clvTable where logs match 'cities'"
	mgram[6] = "select * from clvTable where logs match '/images/cup/'"
	mgram[7] = "select * from clvTable where logs match 'space.gif'"
	mgram[8] = "select * from clvTable where logs match '/cup/15txt'"
	mgram[9] = "select * from clvTable where logs match '11187'"
	mgram[10] = "select * from clvTable where logs match 'french'"
	mgram[11] = "select * from clvTable where logs match 'past_cups/im'"
	mgram[12] = "select * from clvTable where logs match 'ticket_quest'"
	mgram[13] = "select * from clvTable where logs match 'home_tool'"
	mgram[14] = "select * from clvTable where logs match 'past_cups'"
	mgram[15] = "select * from clvTable where logs match 'football.gif'"
	mgram[16] = "select * from clvTable where logs match 'HTTP'"
	mgram[17] = "select * from clvTable where logs match 'comp_bu_stag'"
	mgram[18] = "select * from clvTable where logs match 's102438'"
	mgram[19] = "select * from clvTable where logs match 'paris_stad_'"
	mgram[20] = "select * from clvTable where logs match 'nav_history_'"
	mgram[21] = "select * from clvTable where logs match 's102338.gif'"
	mgram[22] = "select * from clvTable where logs match 'venues'"
	mgram[23] = "select * from clvTable where logs match 'nav_sitemap_'"
	mgram[24] = "select * from clvTable where logs match 'space.gif'"
	mgram[25] = "select * from clvTable where logs match 'hm_nbg.jpg'"
	mgram[26] = "select * from clvTable where logs match 'stage2_brc_'"
	mgram[27] = "select * from clvTable where logs match 'hm_linkf'"
	mgram[28] = "select * from clvTable where logs match 'bg_bottom'"
	mgram[29] = "select * from clvTable where logs match 'cal_paris'"

	for i := 0; i < 30; i++ {
		s := time.Now().UnixMicro()
		q := client.NewQuery(mgram[i], "clv", "")
		if response, err := c.Query(q); err == nil && response.Error() == nil {
			e := time.Now().UnixMicro()
			fmt.Println(len(response.Results))
			fmt.Println("end to end time")
			fmt.Println(float64(e-s) / 1000)
		}
	}*/

	/*mgram := make([]string, 30)
	mgram[0] = "select * from clvTable where logs match '11187'"
	mgram[1] = "select * from clvTable where logs match 'cities'"
	mgram[2] = "select * from clvTable where logs match 'venues'"
	mgram[3] = "select * from clvTable where logs match 'hm_linkf'"
	mgram[4] = "select * from clvTable where logs match 'tickets/'"
	mgram[5] = "select * from clvTable where logs match 'past_cups'"
	mgram[6] = "select * from clvTable where logs match 'space.gif'"
	mgram[7] = "select * from clvTable where logs match 'space.gif'"
	mgram[8] = "select * from clvTable where logs match 'home_tool'"
	mgram[9] = "select * from clvTable where logs match 'bg_bottom'"

	mgram[10] = "select * from clvTable where logs match 'cal_paris'"
	mgram[11] = "select * from clvTable where logs match '/cup/15txt'"
	mgram[12] = "select * from clvTable where logs match 'paris_stad_'"
	mgram[13] = "select * from clvTable where logs match 's102338.gif'"
	mgram[14] = "select * from clvTable where logs match 'stage2_brc_'"
	mgram[15] = "select * from clvTable where logs match 'cal_mont.gif'"
	mgram[16] = "select * from clvTable where logs match '/images/cup/'"
	mgram[17] = "select * from clvTable where logs match 'ticket_quest'"
	mgram[18] = "select * from clvTable where logs match 'nav_history_'"
	mgram[19] = "select * from clvTable where logs match 'comp_bu_stag'"

	mgram[20] = "select * from clvTable where logs match 'past_cups/im'"
	mgram[21] = "select * from clvTable where logs match 's102325.gif HTTP'"
	mgram[22] = "select * from clvTable where logs match 'home_fr_phrase.gif'"
	mgram[23] = "select * from clvTable where logs match 'frntpage.htm HTTP/1.0'"
	mgram[24] = "select * from clvTable where logs match 'ticket_bu_infrance2.gif'"
	mgram[25] = "select * from clvTable where logs match '/english/history/past_cups/'"
	mgram[26] = "select * from clvTable where logs match 'GET /images/home_fr_button.gif HTTP/1.0'"
	mgram[27] = "select * from clvTable where logs match 'GET /english/nav_top_inet.html HTTP/1.0'"
	mgram[28] = "select * from clvTable where logs match 'GET /english/images/nav_sitemap_off.gif'"
	mgram[29] = "select * from clvTable where logs match 'GET /english/playing/images/play_header.gif HTTP/1.0'"

	for i := 0; i < 30; i++ {
		s := time.Now().UnixMicro()
		q := client.NewQuery(mgram[i], "clv", "")
		if response, err := c.Query(q); err == nil && response.Error() == nil {
			e := time.Now().UnixMicro()
			if len(response.Results) == 0 {
				continue
			}
			if len(response.Results[0].Series) == 0 {
				continue
			}
			fmt.Println(len(response.Results[0].Series[0].Values)) //这个都是1
			//fmt.Println(response.Results)
			fmt.Println("end to end time")
			fmt.Println(float64(e-s) / 1000)
		}
	}*/

	/*mtoken := make([]string, 2)
	mtoken[0] = "select * from clvTable where logs match 'GET /english/playing/images/play_header.gif HTTP/1.0'"
	mtoken[1] = "select * from clvTable where logs match 'hm_linkf'"

	for i := 0; i < 2; i++ {
		s := time.Now().UnixMicro()
		q := client.NewQuery(mtoken[i], "clv", "")
		if response, err := c.Query(q); err == nil && response.Error() == nil {
			e := time.Now().UnixMicro()
			fmt.Println(len(response.Results))
			fmt.Println("******************************************")
			fmt.Println("end to end time")
			fmt.Println(float64(e-s) / 1000)
			fmt.Println("******************************************")
		}
	}*/

	/*mtoken := make([]string, 30)
	mtoken[0] = "select * from clvTable where logs match 'GET /french/tickets/images/ticket_bu_infrance2.gif'"
	mtoken[1] = "select * from clvTable where logs match 'GET /english/nav_top_inet.html HTTP/1.0'"
	mtoken[2] = "select * from clvTable where logs match 'GET /images/hm_bg.jpg HTTP/1.0'"
	mtoken[3] = "select * from clvTable where logs match 'ticket_bu_infrance2.gif'"
	mtoken[4] = "select * from clvTable where logs match 'GET /english/images/team_hm_header_shad.gif HTTP/1.0'"
	mtoken[5] = "select * from clvTable where logs match 'GET /images/s102325.gif HTTP/1.0'"
	mtoken[6] = "select * from clvTable where logs match 'GET /english/history/history_of/images/cup'"
	mtoken[7] = "select * from clvTable where logs match 'GET /images/backnews.gif'"
	mtoken[8] = "select * from clvTable where logs match 'GET / HTTP/1.0'"
	mtoken[9] = "select * from clvTable where logs match '11187'"
	mtoken[10] = "select * from clvTable where logs match 'past_cups/images/past_bu_30_off.gif'"
	mtoken[11] = "select * from clvTable where logs match 'nav_tickets_off.gif'"
	mtoken[12] = "select * from clvTable where logs match 'ticket_quest_bg2'"
	mtoken[13] = "select * from clvTable where logs match 'GET /english/images/nav_sitemap_off.gif'"
	mtoken[14] = "select * from clvTable where logs match 'nav_home_off.gif'" //french/images/nav_home_off.gif
	mtoken[15] = "select * from clvTable where logs match 'football.GIF'"
	mtoken[16] = "select * from clvTable where logs match 'GET /images/dot.gif HTTP/1.0'"
	mtoken[17] = "select * from clvTable where logs match 'GET /english/images/nav_hosts_off.gif'"
	mtoken[18] = "select * from clvTable where logs match 'venue_paris_stad_header.gif'"
	mtoken[19] = "select * from clvTable where logs match 'GET /images/ligneb01.gif'"
	mtoken[20] = "select * from clvTable where logs match 'mascot.html'"
	mtoken[21] = "select * from clvTable where logs match 'venues'"
	mtoken[22] = "select * from clvTable where logs match 'index.html'"
	mtoken[23] = "select * from clvTable where logs match 'space.gif'"
	mtoken[24] = "select * from clvTable where logs match 'GET /english/frntpage.htm HTTP/1.0'"
	mtoken[25] = "select * from clvTable where logs match 'comp_stage2_brc_topr.gif'"
	mtoken[26] = "select * from clvTable where logs match 'hm_linkf.gif'"
	mtoken[27] = "select * from clvTable where logs match 'nav_bg_bottom.jpg'"
	mtoken[28] = "select * from clvTable where logs match 'GET /english/images/nav_news_off.gif'"
	mtoken[29] = "select * from clvTable where logs match 'cal_paris.gif'"

	for i := 0; i < 30; i++ {
		s := time.Now().UnixMicro()
		q := client.NewQuery(mtoken[i], "clv", "")
		if response, err := c.Query(q); err == nil && response.Error() == nil {
			e := time.Now().UnixMicro()
			fmt.Println(len(response.Results))
			fmt.Println("******************************************")
			fmt.Println("end to end time")
			fmt.Println(float64(e-s) / 1000)
			fmt.Println("******************************************")
		}
	}*/

	/*mtoken := make([]string, 8)
	mtoken[0] = "select * from clvTable where logs match 'GET /images/hm_bg.jpg HTTP/1.0'"
	mtoken[1] = "select * from clvTable where logs match 'GET /english/images/nav_sitemap_off.gif'"
	mtoken[2] = "select * from clvTable where logs match 'GET /english/teams/teamgroup.htm HTTP/1.0'"
	mtoken[3] = "select * from clvTable where logs match 'GET /english/history/past_cups/images/past_bu_66_off.gif HTTP/1.0'"
	mtoken[4] = "select * from clvTable where logs match 'GET /english/images/team_bu_detail_on.gif HTTP/1.0'"
	mtoken[5] = "select * from clvTable where logs match 'GET /images/32t49807.jpg HTTP/1.0'"
	mtoken[6] = "select * from clvTable where logs match 'GET /english/images/nav_news_off.gif'"
	mtoken[7] = "select * from clvTable where logs match 'GET /french/venues/images/Venue_map_bot_off.gif'"

	for i := 0; i < 8; i++ {
		s := time.Now().UnixMicro()
		q := client.NewQuery(mtoken[i], "clv", "")
		if response, err := c.Query(q); err == nil && response.Error() == nil {
			e := time.Now().UnixMicro()
			fmt.Println(len(response.Results))
			fmt.Println("******************************************")
			fmt.Println("end to end time")
			fmt.Println(float64(e-s) / 1000)
			fmt.Println("******************************************")
		}
	}*/

	/*s := time.Now().UnixMicro()
	q := client.NewQuery("select * from clvTable where logs match 'GET'", "clv", "")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		e := time.Now().UnixMicro()
		fmt.Println(len(response.Results[0].Series))
		fmt.Println("******************************************")
		fmt.Println("end to end time")
		fmt.Println(float64(e-s) / 1000)
		fmt.Println("******************************************")
	}*/

	/*var fuzzyGramQuery = [30]string{
		"select * from clvTable where logs regex 'fran.*'",
		"select * from clvTable where logs regex 'space.*'",
		"select * from clvTable where logs regex 'en.+ish'",
		"select * from clvTable where logs regex 'im.ges'",
		"select * from clvTable where logs regex 's10.327'",
		"select * from clvTable where logs regex 'hm_a.*'",
		"select * from clvTable where logs regex 'his.*ry'",
		"select * from clvTable where logs regex 'GE.*'",
		"select * from clvTable where logs regex 'ticket_b.*'",
		"select * from clvTable where logs regex 'past_c.*ps'",
		"select * from clvTable where logs regex 'past_bu_9.*'",
		"select * from clvTable where logs regex 'dot.*'",
		"select * from clvTable where logs regex '.*fr_phrase.*'",
		"select * from clvTable where logs regex 'team.*io138'",
		"select * from clvTable where logs regex 'hm.*'",
		"select * from clvTable where logs regex 'news.*'",
		"select * from clvTable where logs regex 'play.*'",
		"select * from clvTable where logs regex 's102.*'",
		"select * from clvTable where logs regex 'down.*'",
		"select * from clvTable where logs regex 'team_group.*'",
		"select * from clvTable where logs regex 'GET /fr.*'",
		"select * from clvTable where logs regex 'body.*'",
		"select * from clvTable where logs regex 'lig.*'",
		"select * from clvTable where logs regex 've(nu|ue)es'",
		"select * from clvTable where logs regex 'venue.bu'",
		"select * from clvTable where logs regex 'tic.+et.head'",
		"select * from clvTable where logs regex 'hm_.?f98'",
		"select * from clvTable where logs regex 'venue.*'",
		"select * from clvTable where logs regex 'header.*'",
		"select * from clvTable where logs regex 'f98.top'",
	}

	for i := 0; i < 30; i++ {
		s := time.Now().UnixMicro()
		q := client.NewQuery(fuzzyGramQuery[i], "clv", "")
		if response, err := c.Query(q); err == nil && response.Error() == nil {
			e := time.Now().UnixMicro()
			fmt.Println(len(response.Results))
			fmt.Println("end to end time")
			fmt.Println(float64(e-s) / 1000)
		}
	}*/

	//vtoken 二阶段
	//var fuzzyTokenQuery = [30]string{
	//	"select * from clvTable where logs regex 'lig.*'",
	//	"select * from clvTable where logs regex 'teamgroup.*htm'",
	//	"select * from clvTable where logs regex 'hm_of.icial.gif'",
	//	"select * from clvTable where logs regex 's10.327'",
	//	"select * from clvTable where logs regex 'home_to+l.gif'",
	//	"select * from clvTable where logs regex 'proscrol+.class'",
	//	"select * from clvTable where logs regex 'fr.nch'",
	//	"select * from clvTable where logs regex 'past_c.*ps'",
	//	"select * from clvTable where logs regex 'hm_b.*l.gif'",
	//	"select * from clvTable where logs regex 'prosc.oll.class'",
	//	"select * from clvTable where logs regex '.*fr_phrase.*'",
	//	"select * from clvTable where logs regex 'home.*stars.gif'",
	//	"select * from clvTable where logs regex 'nav_bg_t.p.gif'",
	//	"select * from clvTable where logs regex 'past_c.*s'",
	//	"select * from clvTable where logs regex 'nav.*html'",
	//	"select * from clvTable where logs regex 'lig.*gif'",
	//	"select * from clvTable where logs regex 'hm_of.icial.gif'",
	//	"select * from clvTable where logs regex 's102.*'",
	//	"select * from clvTable where logs regex 'tic.+ets/'",
	//	"select * from clvTable where logs regex 'body.*'",
	//	"select * from clvTable where logs regex 'nav_home.*gif'",
	//	"select * from clvTable where logs regex 'team_group.*gif'",
	//	"select * from clvTable where logs regex 's1023(36|25)'",
	//	"select * from clvTable where logs regex 'nav_(history|tickets)_off.gif'",
	//	"select * from clvTable where logs regex 'news_btn_(kits|letter)_off.gif'",
	//	"select * from clvTable where logs regex 'hm_(day|anime)_e.gif'",
	//	"select * from clvTable where logs regex 'home(_intro.anim.gif|_fr_button.gif)'",
	//	"select * from clvTable where logs regex 'imag.s'",
	//	"select * from clvTable where logs regex '(02how|04luck).gif'",
	//	"select * from clvTable where logs regex 'ligne(b01|01)'",
	//}

	//vgram 二阶段
	//var fuzzyTokenQuery = [30]string{
	//	"select * from clvTable where logs regex 'fran.*'",
	//	"select * from clvTable where logs regex 'space.*'",
	//	"select * from clvTable where logs regex 'en.+ish'",
	//	"select * from clvTable where logs regex 'im.ges'",
	//	"select * from clvTable where logs regex 's10.327'",
	//	"select * from clvTable where logs regex 'hm_a.*'",
	//	"select * from clvTable where logs regex 'his.*ry'",
	//	"select * from clvTable where logs regex 'GE.*'",
	//	"select * from clvTable where logs regex 'ticket_b.*'",
	//	"select * from clvTable where logs regex 'past_c.*ps'",
	//	"select * from clvTable where logs regex 'past_bu_9.*'",
	//	"select * from clvTable where logs regex 'dot.*'",
	//	"select * from clvTable where logs regex '.*fr_phrase.*'",
	//	"select * from clvTable where logs regex 'team.*io138'",
	//	"select * from clvTable where logs regex 'hm.*'",
	//	"select * from clvTable where logs regex 'news.*'",
	//	"select * from clvTable where logs regex 'play.*'",
	//	"select * from clvTable where logs regex 's102.*'",
	//	"select * from clvTable where logs regex 'down.*'",
	//	"select * from clvTable where logs regex 'team_group.*'",
	//	"select * from clvTable where logs regex 'GET /fr.*'",
	//	"select * from clvTable where logs regex 'body.*'",
	//	"select * from clvTable where logs regex 'lig.*'",
	//	"select * from clvTable where logs regex 've(nu|ue)es'",
	//	"select * from clvTable where logs regex 'venue.bu'",
	//	"select * from clvTable where logs regex 'tic.+et.head'",
	//	"select * from clvTable where logs regex 'hm_.?f98'",
	//	"select * from clvTable where logs regex 'venue.*'",
	//	"select * from clvTable where logs regex 'header.*'",
	//	"select * from clvTable where logs regex 'f98.top'",
	//}

	//vtoken 3阶段
	var fuzzyTokenQuery = [30]string{
		"select * from clvTable where logs regex '(02how|04luck).gif'",
		"select * from clvTable where logs regex 'newspr*.htm'",
		"select * from clvTable where logs regex 'archiv(es|ed).gif'",
		"select * from clvTable where logs regex 'infranc*e.*.gif'",
		"select * from clvTable where logs regex 'comp_(hm|mn)_nav.gif'",
		"select * from clvTable where logs regex 'ticket.abroad.header.gif'",
		"select * from clvTable where logs regex 'headtohead(75|86)'",
		"select * from clvTable where logs regex 'cal_lyon?.gif'",
		"select * from clvTable where logs regex 'main(co|te)mp.htm'",
		"select * from clvTable where logs regex '(16)*tradition.gif'",
		"select * from clvTable where logs regex 'marseill(e|s).gif'",
		"select * from clvTable where logs regex 'france(38|37)'",
		"select * from clvTable where logs regex 'montpellier+.gif'",
		"select * from clvTable where logs regex 'player(908|786)'",
		"select * from clvTable where logs regex 'zo+m.htm'",
		"select * from clvTable where logs regex 'comp_bu_calendar*_on.gif'",
		"select * from clvTable where logs regex 'italy(90|8)'",
		"select * from clvTable where logs regex 'read(ing|er)'",
		"select * from clvTable where logs regex '16overlo+k.gif'",
		"select * from clvTable where logs regex 'teambio15(4|5)'",
		"select * from clvTable where logs regex 'foot(ball|s)'",
		"select * from clvTable where logs regex 'beauty+.zip'",
		"select * from clvTable where logs regex 'b(ra|u)cket.gif'",
		"select * from clvTable where logs regex '(15|4)aerial.jpg'",
		"select * from clvTable where logs regex 'dbur(ton|uin).jpg'",
		"select * from clvTable where logs regex 'image+'",
		"select * from clvTable where logs regex 'player(01|10)'",
		"select * from clvTable where logs regex 'frntpag(e|d).htm'",
		"select * from clvTable where logs regex 'us(1703|2345)zp.htm'",
		"select * from clvTable where logs regex 'supplier(s|u)'",
	}

	//vgram 三阶段
	//var fuzzyTokenQuery = [30]string{
	//	"select * from clvTable where logs regex 'past_bu_9.*'",
	//	"select * from clvTable where logs regex 'body.*'",
	//	"select * from clvTable where logs regex 'ticket_b.*'",
	//	"select * from clvTable where logs regex 'dot.*'",
	//	"select * from clvTable where logs regex 'fran.*'",
	//	"select * from clvTable where logs regex '.*fr_phrase.*'",
	//	"select * from clvTable where logs regex '(15|4)aerial'",
	//	"select * from clvTable where logs regex 'teambio15(4|5)'",
	//	//
	//	"select * from clvTable where logs regex 'ac*omodations'",
	//	"select * from clvTable where logs regex 'france(38|37)'",
	//	"select * from clvTable where logs regex 'headtohead(75|86)'",
	//	//
	//	"select * from clvTable where logs regex '.*sea'",
	//	"select * from clvTable where logs regex 'italy(90|8)'",
	//	"select * from clvTable where logs regex 'beauty+'",
	//	"select * from clvTable where logs regex 'zo+m'",
	//	//
	//	"select * from clvTable where logs regex '.*brazilianplayers.*'",
	//	"select * from clvTable where logs regex '(16)*tradition'",
	//	"select * from clvTable where logs regex '16overlo+k'",
	//	"select * from clvTable where logs regex '1*87'",
	//	"select * from clvTable where logs regex 'the+cup'",
	//	"select * from clvTable where logs regex 'box_?saver2'",
	//	"select * from clvTable where logs regex 'links_*on'",
	//	"select * from clvTable where logs regex '32p*49816'",
	//	"select * from clvTable where logs regex '.*tck_pkit_fx_a'",
	//	//
	//	"select * from clvTable where logs regex 'qrimet_?delaunay'",
	//	"select * from clvTable where logs regex '.*latebreak'",
	//	"select * from clvTable where logs regex '/cgi-bin/?trivia'",
	//	"select * from clvTable where logs regex '.*/french/tickets.*images/ticket_bu_?infrance2'",
	//	//
	//	"select * from clvTable where logs regex 'kits?_off'",
	//	"select * from clvTable where logs regex '.*/english/venues/cities/images.*denis/venue_denn_bg'",
	//}
	//
	for i := 0; i < 30; i++ {
		s := time.Now().UnixMicro()
		q := client.NewQuery(fuzzyTokenQuery[i], "clv", "")
		if response, err := c.Query(q); err == nil && response.Error() == nil {
			e := time.Now().UnixMicro()
			if len(response.Results) == 0 {
				continue
			}
			if len(response.Results[0].Series) == 0 {
				continue
			}
			fmt.Println(len(response.Results[0].Series[0].Values))
			fmt.Println("end to end time")
			fmt.Println(float64(e-s) / 1000)
		}
	}

}

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
	mgram[4] = "select * from clvTable where logs match 'sponsor.gif'"
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
	mgram[21] = "select * from clvTable where logs match 'mascot.html'"
	mgram[22] = "select * from clvTable where logs match 'venues'"
	mgram[23] = "select * from clvTable where logs match 'nav_sitemap_'"
	mgram[24] = "select * from clvTable where logs match 'space.gif'"
	mgram[25] = "select * from clvTable where logs match 'frntpage.htm'"
	mgram[26] = "select * from clvTable where logs match 'stage2_brc_'"
	mgram[27] = "select * from clvTable where logs match 'hm_linkf'"
	mgram[28] = "select * from clvTable where logs match 'bg_bottom'"
	mgram[29] = "select * from clvTable where logs match 'cal_paris'"

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

	/*mgram := make([]string, 30)
	mgram[0] = "select * from clvTable where logs match '/cup/15txt'"
	mgram[1] = "select * from clvTable where logs match 'paris_stad_'"
	mgram[2] = "select * from clvTable where logs match 'cities'"
	mgram[3] = "select * from clvTable where logs match 'ticket_quest'"
	mgram[4] = "select * from clvTable where logs match 'comp_bu_stag'"
	mgram[5] = "select * from clvTable where logs match 'ticket_bu_infrance2.gif'"
	mgram[6] = "select * from clvTable where logs match 'GET /images/hm_bg.jpg'"
	mgram[7] = "select * from clvTable where logs match 'GET /images/home_bg_stars.gif'"
	mgram[8] = "select * from clvTable where logs match 'GET /english/images/team_group_header_d.gif'"
	mgram[9] = "select * from clvTable where logs match 'GET /images/cal_nant.gif'"

	mgram[10] = "select * from clvTable where logs match 'GET /english/ProScroll.class'"
	mgram[11] = "select * from clvTable where logs match '/english/member/images/member_header.jpg'"
	mgram[12] = "select * from clvTable where logs match 'GET /images/home_bg_stars.gif'"
	mgram[13] = "select * from clvTable where logs match 'GET /english/images/today_new.gif'"
	mgram[14] = "select * from clvTable where logs match 'GET /images/s102336.gif'"
	mgram[15] = "select * from clvTable where logs match 'GET /images/bord_d.gif'"
	mgram[16] = "select * from clvTable where logs match 'GET /images/home_fr_button.gif HTTP/1.0'"
	mgram[17] = "select * from clvTable where logs match 'GET /english/nav_top_inet.html HTTP/1.0'"
	mgram[18] = "select * from clvTable where logs match 'GET /english/images/nav_sitemap_off.gif'"
	mgram[19] = "select * from clvTable where logs match 'GET /english/playing/images/play_header.gif HTTP/1.0'"

	mgram[20] = "select * from clvTable where logs match 'GET /french/tickets/images/ticket_quest_'"
	mgram[21] = "select * from clvTable where logs match 'GET /french/images/nav_history_off.gif H'"
	mgram[22] = "select * from clvTable where logs match 'GET /french/images/nav_tickets_off.gif H'"
	mgram[23] = "select * from clvTable where logs match 'GET /french/images/news_btn_letter_off.g'"
	mgram[24] = "select * from clvTable where logs match 'GET /english/playing/mascot/images/logo.'"
	mgram[25] = "select * from clvTable where logs match 'GET /images/home_intro.anim.gif HTTP/1.1'"
	mgram[26] = "select * from clvTable where logs match 'GET /french/images/archives.gif HTTP/1.0'"
	mgram[27] = "select * from clvTable where logs match '/history_of/images/france/history_france'"
	mgram[28] = "select * from clvTable where logs match 'GET /french/images/nav_news_off.gif HTTP/1.0'"
	mgram[29] = "select * from clvTable where logs match 'GET /french/history/images/thecup'"

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
			fmt.Println(len(response.Results[0].Series[0].Values))
			//fmt.Println(response.Results)
			fmt.Println("end to end time")
			fmt.Println(float64(e-s) / 1000)
		}
	}*/

	/*mtoken := make([]string, 30)
	mtoken[0] = "select * from clvTable where logs match 'GET'"
	mtoken[1] = "select * from clvTable where logs match 'GET /english'"
	mtoken[2] = "select * from clvTable where logs match 'GET /english/images'"
	mtoken[3] = "select * from clvTable where logs match 'GET /images'"
	mtoken[4] = "select * from clvTable where logs match 'GET /english/images/team_hm_header_shad.gif HTTP/1.0'"
	mtoken[5] = "select * from clvTable where logs match 'GET /images/s102325.gif HTTP/1.0'"
	mtoken[6] = "select * from clvTable where logs match 'GET /english/history/history_of/images/cup'"
	mtoken[7] = "select * from clvTable where logs match 'images/space.gif'"
	mtoken[8] = "select * from clvTable where logs match 'GET / HTTP/1.0'"
	mtoken[9] = "select * from clvTable where logs match '11187'"
	mtoken[10] = "select * from clvTable where logs match 'french/'"
	mtoken[11] = "select * from clvTable where logs match 'nav_tickets_off.gif'"
	mtoken[12] = "select * from clvTable where logs match 'ticket_quest_bg2'"
	mtoken[13] = "select * from clvTable where logs match 'HTTP/1.1'"
	mtoken[14] = "select * from clvTable where logs match '1.0'"
	mtoken[15] = "select * from clvTable where logs match 'football.GIF'"
	mtoken[16] = "select * from clvTable where logs match 'HTTP'"
	mtoken[17] = "select * from clvTable where logs match 'images'"
	mtoken[18] = "select * from clvTable where logs match 's102438'"
	mtoken[19] = "select * from clvTable where logs match 'venue_paris_stad_header.gif'"
	mtoken[20] = "select * from clvTable where logs match 'nav_history_off.gif'"
	mtoken[21] = "select * from clvTable where logs match 'mascot.html'"
	mtoken[22] = "select * from clvTable where logs match 'venues'"
	mtoken[23] = "select * from clvTable where logs match 'index.html'"
	mtoken[24] = "select * from clvTable where logs match 'space.gif'"
	mtoken[25] = "select * from clvTable where logs match 'GET /english/frntpage.htm HTTP/1.0'"
	mtoken[26] = "select * from clvTable where logs match 'comp_stage2_brc_topr.gif'"
	mtoken[27] = "select * from clvTable where logs match 'hm_linkf.gif'"
	mtoken[28] = "select * from clvTable where logs match 'nav_bg_bottom.jpg'"
	mtoken[29] = "select * from clvTable where logs match 'cal_paris.gif'"

	for i := 0; i < 30; i++ {
		s := time.Now().UnixMicro()
		q := client.NewQuery(mtoken[i], "clv", "")
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

	mtoken := make([]string, 30)
	mtoken[0] = "select * from clvTable where logs match 'GET /images/photo02.gif'"
	mtoken[1] = "select * from clvTable where logs match 'GET /english/playing/images/play_hm_mascot.gif'"
	mtoken[2] = "select * from clvTable where logs match 'GET /images/11104.gif HTTP/1.1'"
	mtoken[3] = "select * from clvTable where logs match 'GET /images/cal_steti.gif HTTP/1.0'"
	mtoken[4] = "select * from clvTable where logs match 'GET /images/base.gif HTTP/1.0'"
	mtoken[5] = "select * from clvTable where logs match 'GET /english/playing/images/anim/mascot_on.gif'"
	mtoken[6] = "select * from clvTable where logs match 'GET /english/playing/download/images/box_saver1.gif'"
	mtoken[7] = "select * from clvTable where logs match 'GET /french/images/hm_official.gif'"
	mtoken[8] = "select * from clvTable where logs match 'past_cups/images/past_bu_30_off.gif'"
	mtoken[9] = "select * from clvTable where logs match 'GET /french/venues/images/venue_bu_acomm_on.gif HTTP/1.0'"

	mtoken[10] = "select * from clvTable where logs match 'GET /french/news/3004bres.htm'"
	mtoken[11] = "select * from clvTable where logs match 'GET / HTTP/1.0'"
	mtoken[12] = "select * from clvTable where logs match 'GET /english/history/past_cups/images/past_bracket_bot.gif'"
	mtoken[13] = "select * from clvTable where logs match 'GET /french/tickets/images/ticket_bu_infrance2.gif'"
	mtoken[14] = "select * from clvTable where logs match 'GET /french/tickets/images/ticket_bu_abroad2.gif'"
	mtoken[15] = "select * from clvTable where logs match 'GET /images/hm_bg.jpg HTTP/1.0'"
	mtoken[16] = "select * from clvTable where logs match 'GET /images/s102325.gif HTTP/1.0'"
	mtoken[17] = "select * from clvTable where logs match 'GET /english/frntpage.htm HTTP/1.0'"
	mtoken[18] = "select * from clvTable where logs match 'GET /english/history/history_of/images/cup'"
	mtoken[19] = "select * from clvTable where logs match 'GET /english/images/team_hm_header_shad.gif HTTP/1.0'"

	mtoken[20] = "select * from clvTable where logs match 'GET /french/venues/body.html HTTP/1.0'"
	mtoken[21] = "select * from clvTable where logs match 'GET /english/images/space.gif HTTP/1.1'"
	mtoken[22] = "select * from clvTable where logs match 'GET /images/hm_linkf.gif HTTP/1.1'"
	mtoken[23] = "select * from clvTable where logs match 'GET /images/11101.gif HTTP/1.0'"
	mtoken[24] = "select * from clvTable where logs match '/english/playing/images/banner2.gif HTTP/1.0'"
	mtoken[25] = "select * from clvTable where logs match 'GET /english/images/fpnewstop.gif HTTP/1.0'"
	mtoken[26] = "select * from clvTable where logs match 'GET /images/bord_stories01.gif HTTP/1.0'"
	mtoken[27] = "select * from clvTable where logs match 'GET /images/dburton.jpg HTTP/1.0'"
	mtoken[28] = "select * from clvTable where logs match 'GET /images/base.gif HTTP/1.0'"
	mtoken[29] = "select * from clvTable where logs match '/english/venues/cities/images/denis/venue_denn_bg.jpg'"

	for i := 0; i < 30; i++ {
		s := time.Now().UnixMicro()
		q := client.NewQuery(mtoken[i], "clv", "")
		if response, err := c.Query(q); err == nil && response.Error() == nil {
			e := time.Now().UnixMicro()
			if len(response.Results) == 0 {
				continue
			}
			if len(response.Results[0].Series) == 0 {
				continue
			}
			fmt.Println(len(response.Results[0].Series[0].Values))
			//fmt.Println(response.Results)
			fmt.Println("end to end time")
			fmt.Println(float64(e-s) / 1000)
		}
	}

	/*mtoken := make([]string, 2)
	mtoken[0] = "select * from clvTable where logs match 'GET /english/images/space.gif HTTP/1.1'"
	mtoken[1] = "select * from clvTable where logs match 'GET /images/hm_linkf.gif HTTP/1.1'"
	for i := 0; i < 2; i++ {
		s := time.Now().UnixMicro()
		q := client.NewQuery(mtoken[i], "clv", "")
		if response, err := c.Query(q); err == nil && response.Error() == nil {
			e := time.Now().UnixMicro()
			if len(response.Results) == 0 {
				continue
			}
			if len(response.Results[0].Series) == 0 {
				continue
			}
			fmt.Println(len(response.Results[0].Series[0].Values))
			//fmt.Println(response.Results)
			fmt.Println("end to end time")
			fmt.Println(float64(e-s) / 1000)
		}
	}*/
}

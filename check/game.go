package check

import (
	"reflect"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type Game struct {
	CheckerBase
}

func (g *Game) GetName() string {
	t := reflect.TypeOf(g)
	return t.Elem().Name()
}

func (c Game) PrintStatus() {
	c.CheckerInfo.PrintStatus(c.GetName())
}

func (g *Game) CheckStock() error {
	g.CheckerInfo.LogCheck()
	url := "https://www.game.co.uk/playstation-5"

	ctx, cancels := SetupBrowserContext(g.Options)
	for _, c := range cancels {
		defer c()
	}

	//var res string
	var stockButtons []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Nodes("#contentPanels3 .sectionButton a", &stockButtons, chromedp.NodeVisible),
	)
	if err != nil {
		g.Errors++
		return err
	}
	for _, sb := range stockButtons {
		if sb.Children[0].NodeValue != "Out of stock" {
			g.CheckerInfo.LogStockSeen(url)
			return nil
		}
	}
	return nil
}

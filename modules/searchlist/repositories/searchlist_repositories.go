package repositories

import (
	// "fmt"
	"github.com/nguitarpb/7-solutions/modules/searchlist/entities"
	"io"
	"net/http"
)

type SearchlistRepo struct {}

func NewSearchlistRepository() entities.SearchlistRepository {
	return &SearchlistRepo{}
}

func (r *SearchlistRepo) SearchListDb() (*string, error) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := string(body)
	// fmt.Println(res)
// 	res := `Magna est spare ribs, quis pastrami bacon et proident short ribs consequat picanha.  Adipisicing tenderloin ut esse shoulder cow voluptate eu spare ribs in pork belly ut reprehenderit salami sirloin.  Consectetur turkey cupidatat, eu pork chop ham hock in corned beef chuck.  Pork chop magna beef ribs, qui ea sint ipsum dolor enim.  In laborum t-bone meatloaf ut cow turducken sausage.  Jowl cillum magna eu rump jerky, fugiat pariatur enim ad dolore laboris.  Qui et tempor pork ut id.

// Laborum chuck jerky leberkas pancetta.  Pariatur chicken ham consectetur laborum aliqua turducken cupim chuck ham hock spare ribs fugiat t-bone velit short loin.  Dolor pancetta porchetta, chicken do eu anim deserunt shank reprehenderit pork.  Ham deserunt ut dolore ham hock sirloin.  Boudin in fatback pastrami chislic, tongue chuck ullamco do adipisicing sed strip steak ex pig.  Voluptate in sint sausage.  In ex irure bacon laboris eiusmod aute.`

	return &res, nil
}

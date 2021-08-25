package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type DBHandler struct {
	DB              *sql.DB //############ need to change to gorm, not sure what the gorm sql db is
	ApiKey          string
	ReadyForTraffic bool
}

var (
	splitText  = regexp.MustCompile(`\s*,\s*|\s,*\s*`)
	stopWords2 = regexp.MustCompile("^(i|me|my|myself|we|our|ours|ourselves|you|your|yours|yourself|yourselves|he|him|his|himself|she|her|hers|herself|it|its|itself|they|them|their|theirs|themselves|what|which|who|whom|this|that|these|those|am|is|are|was|were|be|been|being|have|has|had|having|do|does|did|doing|a|an|the|and|but|if|or|because|as|until|while|of|at|by|for|with|about|against|between|into|through|during|before|after|above|below|to|from|up|down|in|out|on|off|over|under|again|values|further|then|once|here|there|when|where|why|how|all|any|both|each|few|more|most|other|some|such|no|nor|not|only|own|same|so|than|too|very|s|t|can|will|just|don|should|now|0o|0s|3a|3b|3d|6b|6o|a|a1|a2|a3|a4|ab|able|about|above|abst|ac|accordance|according|accordingly|across|act|actually|ad|added|adj|ae|af|affected|affecting|affects|after|afterwards|ag|again|against|ah|ain|ain't|aj|al|all|allow|allows|almost|alone|along|already|also|although|always|am|among|amongst|amoungst|amount|an|and|announce|another|any|anybody|anyhow|anymore|anyone|anything|anyway|anyways|anywhere|ao|ap|apart|apparently|appear|appreciate|appropriate|approximately|ar|are|aren|arent|aren't|arise|around|as|a's|aside|ask|asking|associated|at|au|auth|av|available|aw|away|awfully|ax|ay|az|b|b1|b2|b3|ba|back|bc|bd|be|became|because|become|becomes|becoming|been|before|beforehand|begin|beginning|beginnings|begins|behind|being|believe|below|beside|besides|best|better|between|beyond|bi|bill|biol|bj|bk|bl|bn|both|bottom|bp|br|brief|briefly|bs|bt|bu|but|bx|by|c|c1|c2|c3|ca|call|came|can|cannot|cant|can't|cause|causes|cc|cd|ce|certain|certainly|cf|cg|ch|changes|ci|cit|cj|cl|clearly|cm|c'mon|cn|co|com|come|comes|con|concerning|consequently|consider|considering|contain|containing|contains|corresponding|could|couldn|couldnt|couldn't|course|cp|cq|cr|cry|cs|c's|ct|cu|currently|cv|cx|cy|cz|d|d2|da|date|dc|dd|de|definitely|describe|described|despite|detail|df|di|did|didn|didn't|different|dj|dk|dl|do|does|doesn|doesn't|doing|don|done|don't|down|downwards|dp|dr|ds|dt|du|due|during|dx|dy|e|e2|e3|ea|each|ec|ed|edu|ee|ef|effect|eg|ei|eight|eighty|either|ej|el|eleven|else|elsewhere|em|empty|en|end|ending|enough|entirely|eo|ep|eq|er|es|especially|est|et|et-al|etc|eu|ev|even|ever|every|everybody|everyone|everything|everywhere|ex|exactly|example|except|ey|f|f2|fa|far|fc|few|ff|fi|fifteen|fifth|fify|fill|find|fire|first|five|fix|fj|fl|fn|fo|followed|following|follows|for|former|formerly|forth|forty|found|four|fr|from|front|fs|ft|fu|full|further|furthermore|fy|g|ga|gave|ge|get|gets|getting|gi|give|given|gives|giving|gj|gl|go|goes|going|gone|got|gotten|gr|greetings|gs|gy|h|h2|h3|had|hadn|hadn't|happens|hardly|has|hasn|hasnt|hasn't|have|haven|haven't|having|he|hed|he'd|he'll|hello|help|hence|her|here|hereafter|hereby|herein|heres|here's|hereupon|hers|herself|hes|he's|hh|hi|hid|him|himself|his|hither|hj|ho|home|hopefully|how|howbeit|however|how's|hr|hs|http|hu|hundred|hy|i|i2|i3|i4|i6|i7|i8|ia|ib|ibid|ic|id|i'd|ie|if|ig|ignored|ih|ii|ij|il|i'll|im|i'm|immediate|immediately|importance|important|in|inasmuch|inc|indeed|index|indicate|indicated|indicates|information|inner|insofar|instead|interest|into|invention|inward|io|ip|iq|ir|is|isn|isn't|it|itd|it'd|it'll|its|it's|itself|iv|i've|ix|iy|iz|j|jj|jr|js|jt|ju|just|k|ke|keep|keeps|kept|kg|kj|km|know|known|knows|ko|l|l2|la|largely|last|lately|later|latter|latterly|lb|lc|le|least|les|less|lest|let|lets|let's|lf|like|liked|likely|line|little|lj|ll|ll|ln|lo|look|looking|looks|los|lr|ls|lt|ltd|m|m2|ma|made|mainly|make|makes|many|may|maybe|me|mean|means|meantime|meanwhile|merely|mg|might|mightn|mightn't|mill|million|mine|miss|ml|mn|mo|more|moreover|most|mostly|move|mr|mrs|ms|mt|mu|much|mug|must|mustn|mustn't|my|myself|n|n2|na|name|namely|nay|nc|nd|ne|near|nearly|necessarily|necessary|need|needn|needn't|needs|neither|never|nevertheless|new|next|ng|ni|nine|ninety|nj|nl|nn|no|nobody|non|none|nonetheless|noone|nor|normally|nos|not|noted|nothing|novel|nowhere|nr|ns|nt|ny|o|oa|ob|obtain|obtained|obviously|oc|od|of|off|often|og|oh|oi|oj|ok|okay|ol|old|om|omitted|on|once|one|ones|only|onto|oo|op|oq|or|ord|os|ot|other|others|otherwise|ou|ought|our|ours|ourselves|out|outside|over|overall|ow|owing|own|ox|oz|p|p1|p2|p3|page|pagecount|pages|par|part|particular|particularly|pas|past|pc|pd|pe|per|perhaps|pf|ph|pi|pj|pk|pl|placed|please|plus|pm|pn|po|poorly|possible|possibly|potentially|pp|pq|pr|predominantly|present|presumably|previously|primarily|probably|promptly|proud|provides|ps|pt|pu|put|py|q|qj|qu|que|quickly|quite|qv|r|r2|ra|ran|rather|rc|rd|re|readily|really|reasonably|recent|recently|ref|refs|regarding|regardless|regards|related|relatively|research|research-articl|respectively|resulted|resulting|results|rf|rh|ri|right|rj|rl|rm|rn|ro|rq|rr|rs|rt|ru|run|rv|ry|s|s2|sa|said|same|saw|say|saying|says|sc|sd|se|sec|second|secondly|section|see|seeing|seem|seemed|seeming|seems|seen|self|selves|sensible|sent|serious|seriously|seven|several|sf|shall|shan|shan't|she|shed|she'd|she'll|shes|she's|should|shouldn|shouldn't|should've|show|showed|shown|showns|shows|si|side|significant|significantly|similar|similarly|since|sincere|six|sixty|sj|sl|slightly|sm|sn|so|some|somebody|somehow|someone|somethan|something|sometime|sometimes|somewhat|somewhere|soon|sorry|sp|specifically|specified|specify|specifying|sq|sr|ss|st|still|stop|strongly|sub|substantially|successfully|such|sufficiently|suggest|sup|sure|sy|system|sz|t|t1|t2|t3|take|taken|taking|tb|tc|td|te|tell|ten|tends|tf|th|than|thank|thanks|thanx|that|that'll|thats|that's|that've|the|their|theirs|them|themselves|then|thence|there|thereafter|thereby|thered|therefore|therein|there'll|thereof|therere|theres|there's|thereto|thereupon|there've|these|they|theyd|they'd|they'll|theyre|they're|they've|thickv|thin|think|third|this|thorough|thoroughly|those|thou|though|thoughh|thousand|three|throug|through|throughout|thru|thus|ti|til|tip|tj|tl|tm|tn|to|together|too|took|top|toward|towards|tp|tq|tr|tried|tries|truly|try|trying|ts|t's|tt|tv|twelve|twenty|twice|two|tx|u|u201d|ue|ui|uj|uk|um|un|under|unfortunately|unless|unlike|unlikely|until|unto|uo|up|upon|ups|ur|us|use|used|useful|usefully|usefulness|uses|using|usually|ut|v|va|value|various|vd|ve|ve|very|via|viz|vj|vo|vol|vols|volumtype|vq|vs|vt|vu|w|wa|want|wants|was|wasn|wasnt|wasn't|way|we|wed|we'd|welcome|well|we'll|well-b|went|were|we're|weren|werent|weren't|we've|what|whatever|what'll|whats|what's|when|whence|whenever|when's|where|whereafter|whereas|whereby|wherein|wheres|where's|whereupon|wherever|whether|which|while|whim|whither|who|whod|whoever|whole|who'll|whom|whomever|whos|who's|whose|why|why's|wi|widely|will|willing|wish|with|within|without|wo|won|wonder|wont|won't|words|world|would|wouldn|wouldnt|wouldn't|www|x|x1|x2|x3|xf|xi|xj|xk|xl|xn|xo|xs|xt|xv|xx|y|y2|yes|yet|yj|yl|you|youd|you'd|you'll|your|youre|you're|yours|yourself|yourselves|you've|yr|ys|yt|z|zero|zi|zz)$")
)

func newResponse(c echo.Context, msg string, resBool string, httpStatus int, data *[]interface{}) error {
	responseJson := struct {
		Msg     string //message
		ResBool string //boolean response
		Data    []interface{}
	}{
		msg,
		resBool,
		*data,
	}

	return c.JSON(httpStatus, responseJson) // encode to json and send
}

// //checks the api key
// func checkKey(c echo.Context, dbHandler *DBHandler) error {
// 	//############ need to sub dbHandler into the appropriate struct to access api key ##############
// 	// check if key matches
// 	if c.QueryParam("key") != dbHandler.ApiKey { ////
// 		// if api key does not match
// 		newResponse(c, "Forbidden", "false", http.StatusForbidden, nil)
// 		return errors.New("incorrect api key supplied")
// 	}
// 	return nil
// }

// for client to check if api is active
func HealthCheckLiveness(c echo.Context) error {
	return newResponseSimple(c, "nil", "true", http.StatusOK)
}


// Opens db and returns a struct to access it
func OpenDB() *DBHandler {

	// check environment for the database url
	err := godotenv.Load("go.env")
	if err != nil {
		panic(err.Error())
	}

	databaseURL := os.Getenv("DATABASE_URL_MYSQL")

	//load database connection
	db, err1 := sql.Open("mysql", databaseURL)

	if err1 != nil {
		panic(err.Error())
	} else {
		fmt.Println("no issue")
	}

	dbHandler := DBHandler{db, "1", false}
	return &dbHandler
}

//Retrieve one Restaurants
func (dbHandler *DBHandler) GetRestaurant(c echo.Context) error {

	//get id param
	id := c.QueryParam("ID")

	restaurant := models.Restaurant
	results, err1 := DBHandlerMysql.DB.Query("Select * FROM Restaurant WHERE ID = ?", id)
	defer results.Close()

	if err1 != nil {
		fmt.Println(err1.Error())
		return newResponse(c, "Bad Request", "false", http.StatusBadRequest, nil)
	}

	//scan mysql result
	err2 := results.Scan(&restaurant.ID,
		restaurant.Name,
		restaurant.Description,
		restaurant.Address,
		restaurant.PostalCode)

	if err2 != nil {
		fmt.Println(err2.Error())
		return newResponse(c, "Bad Request", "false", http.StatusBadRequest, nil)
	}

	//return json
	return newResponse(c, "Bad Request", "false", http.StatusBadRequest, &[]interface{}{restaurant})

}

//Retrieve All Restaurants
func (dbHandler *DBHandler) GetRestaurantAll(c echo.Context) error {

	//get param id
	id := c.QueryParam("ID")

	results, err1 := DBHandlerMysql.DB.Query("Select * FROM Restaurant", id)
	defer results.Close()

	if err1 != nil {
		fmt.Println(err1.Error())
		return newResponse(c, "Bad Request", "false", http.StatusBadRequest, nil)
	}

	//gets mysql rows and put into interface array
	restaurantArr := []interface{}{}
	for results.Next() {
		restaurant := models.Restaurant
		err := results.Scan(&restaurant.ID,
			restaurant.Name,
			restaurant.Description,
			restaurant.Address,
			restaurant.PostalCode)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			restaurantArr = append(restaurantArr, restaurant)
		}

	return newResponse(c, "Bad Request", "false", http.StatusBadRequest, &restaurantArr)
}

//Return Restaurants based on search
func (dbHandler *DBHandler) SearchRestaurant(c echo.Context) error {
	searchTerm := c.QueryParam("search")

	scoreArr := []int{}
	sortIndexArr := []int{}
	returnArr := []interface{}{}

	searchTermSplit := CleanWord(searchTerm, splitText, stopWords2)

	if searchTerm == "restaurant" {

		//search through resstaurant
		results, err := DBHandlerMysql.DB.Query("Select * FROM Restaurant")
		defer results.Close()

		if err != nil { //err means username not found, ok to proceed
			return newResponse(c, "Bad Request", "false", http.StatusBadRequest, nil)
		}

		i := 0

		for results.Next() {
			restaurant := models.Restaurant
			err = results.Scan(&restaurant.ID,
				restaurant.Name,
				restaurant.Description,
				restaurant.Address,
				restaurant.PostalCode)

			if err != nil {
				fmt.Println(err.Error())

			} else {
				score := searchItem(searchTerm, searchTermSplit, restaurant.Name, restaurant.Description)

				scoreArr = append(scoreArr, score)
				sortIndexArr = append(sortIndexArr, i)
				returnArr = append(returnArr, restaurant)
				i++
			}
		}

	} else {

		results, err := DBHandlerMysql.DB.Query("Select * FROM Food")
		defer results.Close()

		if err != nil { //err means username not found, ok to proceed
			return newResponse(c, "Bad Request", "false", http.StatusBadRequest, nil)
		}

		i := 0

		for results.Next() {
			food := models.Food
			err = results.Scan(&food.ID,
				&food.Name,
				food.ShopID,
				food.Calories,
				food.Sugary,
				food.Description,
				food.Halal,
				food.Vegan)

			if err != nil {
				fmt.Println(err.Error())

			} else {
				score := searchItem(searchTerm, searchTermSplit, food.Name, food.Description)

				scoreArr = append(scoreArr, score)
				sortIndexArr = append(sortIndexArr, i)
				returnArr = append(returnArr, food)
				i++
			}
		}
	}

	// sort the items by score
	_, sortArr2 := MergeSort(scoreArr, sortIndexArr)
	maxLen := len(scoreArr)
	returnArrSorted := []interface{}{}

	for idx := maxLen - 1; idx >= 0; idx-- { //sorts results in descending order
		newRow := returnArr[sortArr2[idx]]
		returnArrSorted = append(returnArrSorted, newRow)
		// sortInd = append(sortInd, newRow.(model.ItemListing).ID)
	}

	// variable not used, but might be in the future

	return newResponse(c, "Bad Request", "false", http.StatusBadRequest, &returnArrSorted)
}

func searchItem(searchTerm string, searchTermSplit []string, itemName string, ItemDesc string) int {
	points := 0

	// check if search term appears in item name
	if strings.Contains(itemName, searchTerm) {
		points += 5
	}

	// check if search term appears in description
	if strings.Contains(ItemDesc, searchTerm) {
		points += 3
	}

	//continue search if the search term splits into 2 or more words, like nasi lemak
	if len(searchTermSplit) > 1 {
		for _, word := range searchTermSplit {
			if strings.Contains(itemName, word) {
				points += 2
			}
			if strings.Contains(ItemDesc, word) {
				points += 1
			}
		}
	}
	return points
}

//Retrieve All Food from restaurant
func (dbHandler *DBHandler) GetFoodShopID(c echo.Context) error {


	id := c.QueryParam("ID")
	// result := db.Find(&users) // get all


	results, err1 := DBHandlerMysql.DB.Query("Select * FROM Food WHERE ShopID = ?", id)
	if err1 != nil {
		fmt.Println(err1.Error())
		return newResponse(c, "Bad Request", "false", http.StatusBadRequest, nil)
	}

	defer results.Close()

	foodArr := []interface{}{}
	for results.Next() {
		food := models.Food
		err = results.Scan(&food.ID,
			&food.Name,
			food.ShopID,
			food.Calories,
			food.Sugary,
			food.Description,
			food.Halal,
			food.Vegan)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			foodArr = append(foodArr, food)
		}

	return newResponse(c, "Bad Request", "false", http.StatusBadRequest, &foodArr)
}

func InsertSort(arr []int, arrSort []int) ([]int, []int) {
	len1 := len(arr)
	for i := 1; i < len1; i++ {
		temp1 := arr[i]
		tempSort := arrSort[i]
		i2 := i
		for i2 > 0 && arr[i2-1] > temp1 {
			arr[i2] = arr[i2-1]
			arrSort[i2] = arrSort[i2-1]
			i2--
		}
		arr[i2] = temp1
		arrSort[i2] = tempSort
	}
	// fmt.Println(arr, arrSort)
	return arr, arrSort
}

// arrSort is the index
// arr is the arr to be sorted
func MergeSort(arr []int, arrSort []int) ([]int, []int) {
	len1 := int(len(arr))
	len2 := int(len1 / 2)
	if len1 <= 5 {
		return InsertSort(arr, arrSort)
	} else {
		arr1, arrSort1 := MergeSort(arr[len2:], arrSort[len2:])
		arr2, arrSort2 := MergeSort(arr[:len2], arrSort[:len2])
		tempArr := make([]int, len1)
		tempArrSort := make([]int, len1)
		i := 0
		for len(arr1) > 0 && len(arr2) > 0 {
			if arr1[0] < arr2[0] {
				tempArr[i] = arr1[0]
				tempArrSort[i] = arrSort1[0]
				arr1 = arr1[1:]
				arrSort1 = arrSort1[1:]
			} else {
				tempArr[i] = arr2[0]
				tempArrSort[i] = arrSort2[0]
				arr2 = arr2[1:]
				arrSort2 = arrSort2[1:]
			}
			i++
		}
		if len(arr1) == 0 {
			for j := 0; j < len(arr2); j++ {
				// fmt.Println(j, len(arr2), arr2, arr1, tempArr)
				tempArr[i] = arr2[j]
				tempArrSort[i] = arrSort2[j]
				i++
			}
		} else {
			for j := 0; j < len(arr1); j++ {
				tempArr[i] = arr1[j]
				tempArrSort[i] = arrSort1[j]
				i++
			}
		}
		return tempArr, tempArrSort
	}
}

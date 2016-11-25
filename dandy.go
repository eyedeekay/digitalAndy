package main

import "bufio"
import "fmt"
import "flag"
import "image"
import "image/color"
import "image/png"
import "os"
import "math"
import "math/rand"
import "strconv"
import "strings"
import "time"

func readConfig(config string)(error, []string){
	var result []string
	var err error
	fmt.Printf("%v\n", config)
	file, err := os.Open(config)
        if err != nil {
                return err, nil
        }
	defer file.Close()
	scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                result = append(result, scanner.Text())
        }
	fmt.Printf("%v", err)
	return err, result
	
}

func splitString(line string)([]string){
	var presult []string
	var result []string
	presult = strings.Split(line, ";")
	for _, word := range presult {
		result = append(result, strings.Replace(word, ";", "", -1))
	}
	return result
}

func splitWord(word string)([]string){
	var presult []string
	var result []string
	presult0 := strings.TrimSpace(word)
	presult = strings.Split(presult0, " ")
	for _, word := range presult {
		result = append(result, strings.Replace(word, " ", "", -1))
	}
	return result
}

func loadSkeleton(path string)(error, []string){
	var result []string
	var err error
	//fmt.Printf("%v", path)
	if strings.Contains(path, "path=") {
		cleaned := strings.Replace(path, "path=", "", -1)
		cleaned = strings.Replace(cleaned, ";", "", -1)
		err, result = readConfig(cleaned)
	}else if strings.Contains(path, "Path=") {
		cleaned := strings.Replace(path, "Path=", "", -1)
		cleaned = strings.Replace(cleaned, ";", "", -1)
		err, result = readConfig(cleaned)
	}else{
		return err, nil
	}
	return err, result
}

func loadColor(element string)([]string){
	var loaded []string
	if strings.Contains(element, "lcolor") {
		cleaned := strings.Replace(element, "lcolor", "", -1)
		presult := strings.TrimSpace(cleaned)
		fmt.Printf("%v", "loading color : ")
		fmt.Printf("%v\n", presult)
		loaded = splitWord(presult)
	}else if strings.Contains(element, "Lcolor") {
		cleaned := strings.Replace(element, "Lcolor", "", -1)
		presult := strings.TrimSpace(cleaned)
		fmt.Printf("%v", "loading color : ")
		fmt.Printf("%v\n", presult)
		loaded = splitWord(presult)
	}else{
		return loaded
	}
	return loaded
}

func randomColor(colorwords []string, colorlist [][]string)([]string){
	var szcw = len(colorwords)
	var szcl = len(colorlist)
	fmt.Printf("%v", "color palette : ")
	fmt.Printf("%v", szcl)
	var result []string
        r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
        var rcolor int = r.Intn(szcw)
	for _, item := range colorlist {
		if len(item) > 0 {
			if item[0] == colorwords[rcolor] {
				result = item
				fmt.Printf("%v", "selecting color : " )
				fmt.Printf("%v", colorwords[rcolor])
				break
			}
		}
	}
	return result
}

func formatColor(picked []string)([]uint8){
	var result []uint8
	fmt.Printf("%v", "RGBT : ")
	for _, i := range picked {
		fmt.Printf("%v", " : ")
		fmt.Printf("%v", i)
	}
	fmt.Printf("%v\n", " ")
	var pR,_ = strconv.ParseUint(picked[1],10,8)
	result = append(result, uint8(pR))
	var pG,_ = strconv.ParseUint(picked[2],10,8)
	result = append(result, uint8(pG))
	var pB,_ = strconv.ParseUint(picked[3],10,8)
	result = append(result, uint8(pB))
	var pT,_ = strconv.ParseUint(picked[4],10,8)
	result = append(result, uint8(pT))
	return result
}

func selectColor(element string, colorlist [][]string)([]uint8){
	var elect []uint8
	if strings.Contains(element, "scolor") {
		cleaned := strings.Replace(element, "scolor", "", -1)
		colorWords := splitWord(cleaned)
		fmt.Printf("%v", "color selections : ")
		fmt.Printf("%v", cleaned)
		elect = formatColor(randomColor(colorWords, colorlist))
	}else if strings.Contains(element, "Scolor") {
		cleaned := strings.Replace(element, "Scolor", "", -1)
		colorWords := splitWord(cleaned)
		fmt.Printf("%v", "color selections : ")
		fmt.Printf("%v", cleaned)
		elect = formatColor(randomColor(colorWords, colorlist))
	}
	return elect
}

func selectPoint(element string)([]int){
	//fmt.Printf("%v\n", element)
	var pointCoords []string
	var coords []int
	if strings.Contains(element, "point") {
		cleaned := strings.Replace(element, "point", "", -1)
		pointCoords = splitWord(cleaned)
		X,_ := strconv.ParseInt(pointCoords[0], 10, 8)
		coords = append(coords, int(X))
		Y,_ := strconv.ParseInt(pointCoords[1], 10, 8)
		coords = append(coords, int(Y))
	}else if strings.Contains(element, "Point") {
		cleaned := strings.Replace(element, "Point", "", -1)
		pointCoords = splitWord(cleaned)
		X,_ := strconv.ParseInt(pointCoords[0], 10, 8)
		coords = append(coords, int(X))
		Y,_ := strconv.ParseInt(pointCoords[1], 10, 8)
		coords = append(coords, int(Y))
	}
	return coords
}

func selectRect(element string)([]int){
	var pointCoords []string
	var coords []int
	if strings.Contains(element, "rect") {
		cleaned := strings.Replace(element, "rect", "", -1)
		pointCoords = splitWord(cleaned)
		pX,_ := strconv.ParseInt(pointCoords[0], 10, 8)
		pY,_ := strconv.ParseInt(pointCoords[1], 10, 8)
		pW,_ := strconv.ParseInt(pointCoords[2], 10, 8)
		pH,_ := strconv.ParseInt(pointCoords[3], 10, 8)
		for Y := int(pY); Y < int(pY) + int(pH); Y++ {
			for X := int(pX); X < int(pX) + int(pW); X++ {
				coords = append(coords, X)
				coords = append(coords, Y)
			}
		}
	}else if strings.Contains(element, "Rect") {
		cleaned := strings.Replace(element, "Rect", "", -1)
		pointCoords = splitWord(cleaned)
		pX,_ := strconv.ParseInt(pointCoords[0], 10, 8)
		pY,_ := strconv.ParseInt(pointCoords[1], 10, 8)
		pW,_ := strconv.ParseInt(pointCoords[2], 10, 8)
		pH,_ := strconv.ParseInt(pointCoords[3], 10, 8)
		for Y := int(pY); Y < int(pY) + int(pH); Y++ {
			for X := int(pX); X < int(pX) + int(pW); X++ {
				coords = append(coords, X)
				coords = append(coords, Y)
			}
		}
	}
	return coords
}

func Round(val int) (int) {
	var round float64
	pow := math.Pow(10, float64(0))
	digit := pow * float64(val)
	_, div := math.Modf(digit)
	if div >= float64(.5) {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal := round / pow
	return int(newVal)
}

func selectRound(element string)([]int){
	var pointCoords []string
	var coords []int
	if strings.Contains(element, "round") {
		cleaned := strings.Replace(element, "round", "", -1)
		pointCoords = splitWord(cleaned)
		pX,_ := strconv.ParseInt(pointCoords[0], 10, 8)
		pY,_ := strconv.ParseInt(pointCoords[1], 10, 8)
		pW,_ := strconv.ParseInt(pointCoords[2], 10, 8)
		pH,_ := strconv.ParseInt(pointCoords[3], 10, 8)
		for Y := int(pY) + Round(int(pY) / 2); Y < int(pY) + (int(pH) / 2); Y++ {
			for X := int(pX) + Round(int(pX) / 2); X < int(pX) + (int(pW) / 2); X++ {
				coords = append(coords, X)
				coords = append(coords, Y)
			}
		}
		for Y := int(pY); Y < int(pY) + Round(int(pH) / 3); Y++ {
			for X := int(pX); X < int(pX) + Round(int(pW) / 3); X++ {
				coords = append(coords, X)
				coords = append(coords, Y)
			}
		}
	}else if strings.Contains(element, "Round") {
		cleaned := strings.Replace(element, "Round", "", -1)
		pointCoords = splitWord(cleaned)
		pX,_ := strconv.ParseInt(pointCoords[0], 10, 8)
		pY,_ := strconv.ParseInt(pointCoords[1], 10, 8)
		pW,_ := strconv.ParseInt(pointCoords[2], 10, 8)
		pH,_ := strconv.ParseInt(pointCoords[3], 10, 8)
		for Y := int(pY) + Round(int(pY) / 2); Y < int(pY) + Round(int(pH) / 2); Y++ {
			for X := int(pX) + Round(int(pX) / 2); X < int(pX) + Round(int(pW) / 2); X++ {
				coords = append(coords, X)
				coords = append(coords, Y)
			}
		}
		for Y := int(pY); Y < int(pY) + Round(int(pH) / 3); Y++ {
			for X := int(pX); X < int(pX) + Round(int(pW) / 3); X++ {
				coords = append(coords, X)
				coords = append(coords, Y)
			}
		}
	}
	return coords
}

func runSkeleton(skel string, colorlist *[][]string)(error){
	err, arr := readConfig(skel)
	//fmt.Printf("%v\n", skel)
	loopConfigs(err, arr, colorlist)
	return err
}

func runLine(line string, colorlist *[][]string)([]uint8, [][]int){
	//*colorlist = append(*colorlist, loadColor(line))
	var coordResult [][]int
	var colorResult []uint8
	seg := splitString(line)
	for _, word := range seg {
		fmt.Printf("%v\n", word)
		la := loadColor(word)
		if len(la) > 0 {
			*colorlist = append(*colorlist, la)
		}
		lc := selectColor(word, *colorlist)
		if len(lc) > 0 {
			colorResult = lc
		}
		lp := selectPoint(word)
		if len(lp) > 0 {
			coordResult = append(coordResult, lp)
		}
		lr := selectRect(word)
		if len(lr) > 0 {
			coordResult = append(coordResult, lr)
		}
		lo := selectRound(word)
		if len(lo) > 0 {
			coordResult = append(coordResult, lo)
		}
	}
	return colorResult, coordResult
}

func RandStringBytes() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func loopConfigs(err error, strlist []string, colorlist *[][]string)(int){
	img := image.NewRGBA(image.Rect(0, 0, 32, 32))
	var sz = len(strlist)
	var result int
	for i, item := range strlist {
		//fmt.Printf("%v\n", item)
		_, temp := loadSkeleton(item)
		if temp != nil {
			for _,i := range temp {
				c,p := runLine(i, colorlist)
				fmt.Printf("%v", p)
				fmt.Printf("%v\n", c)
				if len(p) > 0 {
					for _, point := range p {
						if len(point) > 0 {
							if len(c) > 0 {
								fmt.Printf("%v", "imageset")
								img.Set(point[0], point[1], color.RGBA{c[0], c[1], c[2], c[3]})
							}
						}
					}
				}
			}
		}else{
			//fmt.Printf("%v\n", item)
			_ = runSkeleton(item, colorlist)
		}
		result = sz - i
	}
	var name =  RandStringBytes()+ ".png"
        f, _ := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
        defer f.Close()
        png.Encode(f, img)
	return result
}

func main(){
	opts := flag.String("conf","config.txg","path to a configuration file")
	confErr, configArray := readConfig(*opts)
	var colors [][]string
	loopConfigs(confErr, configArray, &colors)
}
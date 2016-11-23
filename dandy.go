package main

import "bufio"
//import "fmt"
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
	file, err := os.Open(config)
        if err != nil {
                return err, nil
        }
	defer file.Close()
	scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                result = append(result, scanner.Text())
        }
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
	presult = strings.Split(word, " ")
	for _, word := range presult {
		result = append(result, strings.Replace(word, " ", "", -1))
	}
	return result
}

func loadSkeleton(path string)(error, []string){
	var result []string
	var err error
	if strings.Contains(path, "path=") {
		err, result = readConfig(strings.Replace(path, "path=", "", -1))
	}else if strings.Contains(path, "Path=") {
		err, result = readConfig(strings.Replace(path, "Path=", "", -1))
	}else{
		return err, nil
	}
	return err, result
}

func loadColor(element string)([]string){
	var loaded []string
	if strings.Contains(element, "color") {
		cleaned := strings.Replace(element, "color", "", -1)
		loaded = splitWord(cleaned)
	}else if strings.Contains(element, "Color") {
		cleaned := strings.Replace(element, "color", "", -1)
		loaded = splitWord(cleaned)
	}else{
		return loaded
	}
	return loaded
}

func randomColor(colorwords []string, colorlist [][]string)([]string){
	var szcw = len(colorwords)
	//var szcl = len(colorlist)
	var result []string
        r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
        var rcolor int = r.Intn(szcw)
	for _, item := range colorlist {
		if item[0] == colorwords[rcolor]{
			result = item
		}
	}
	return result
}

func formatColor(picked []string)([]uint8){
	var result []uint8
	if strings.Contains(picked[1], "R") {
		var pR,_ = strconv.ParseInt(strings.Replace(picked[1], "R ", "", -1), 10, 8)
		result = append(result, uint8(pR))
	}else if strings.Contains(picked[2], "G") {
		var pG,_ = strconv.ParseInt(strings.Replace(picked[2], "G ", "", -1), 10, 8)
		result = append(result, uint8(pG))
	}else if strings.Contains(picked[3], "B") {
		var pB,_ = strconv.ParseInt(strings.Replace(picked[3], "B ", "", -1), 10, 8)
		result = append(result, uint8(pB))
	}else if strings.Contains(picked[4], "T") {
		var pT,_ = strconv.ParseInt(strings.Replace(picked[4], "T ", "", -1), 10, 8)
		result = append(result, uint8(pT))
	}
	return result
}

func selectColor(element string, colorlist [][]string)([]uint8){
	var elect []uint8
	if strings.Contains(element, "colors") {
		cleaned := strings.Replace(element, "colors", "", -1)
		colorWords := splitWord(cleaned)
		elect = formatColor(randomColor(colorWords, colorlist))
	}else if strings.Contains(element, "Colors") {
		cleaned := strings.Replace(element, "Colors", "", -1)
		colorWords := splitWord(cleaned)
		elect = formatColor(randomColor(colorWords, colorlist))
	}
	return elect
}

func selectPoint(element string)([]int){
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
	if strings.Contains(element, "point") {
		cleaned := strings.Replace(element, "point", "", -1)
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
	}else if strings.Contains(element, "Point") {
		cleaned := strings.Replace(element, "Point", "", -1)
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
	if strings.Contains(element, "point") {
		cleaned := strings.Replace(element, "point", "", -1)
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
	}else if strings.Contains(element, "Point") {
		cleaned := strings.Replace(element, "Point", "", -1)
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

func runSkeleton(syn string, colorlist [][]string)(error){
	err, arr := readConfig(syn)
	loopConfigs(err, arr, colorlist)
	return err
}

func runLine(line string, colorlist [][]string)([]uint8, [][]int){
	colorlist = append(colorlist, loadColor(line))
	var coordResult [][]int
	colorResult := selectColor(line, colorlist)
	coordResult = append(coordResult, selectPoint(line))
	coordResult = append(coordResult, selectRect(line))
	coordResult = append(coordResult, selectRound(line))
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

func loopConfigs(err error, strlist []string, colorlist [][]string)(int){
	img := image.NewRGBA(image.Rect(0, 0, 32, 32))
	var sz = len(strlist)
	var result int
	for i, item := range strlist {
		element := splitString(item)
		for _, word := range element {
			_, temp := loadSkeleton(word)
			if temp != nil {
				_ = runSkeleton(word, colorlist)
			}else{
				c,p := runLine(word, colorlist)
				for _, point := range p {
					img.Set(point[0], point[1], color.RGBA{c[0], c[1], c[2], c[3]})
				}
			}
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
	loopConfigs(confErr, configArray, colors)
	
}
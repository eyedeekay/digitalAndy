package main

import "bufio"
import "fmt"
import "image"
import "image/color"
import "image/png"
import "os"
import "math/rand"
import "strconv"
import s "strings"
import "time"

func readLines(path string) ([]string, error) {
        file, err := os.Open(path)
        if err != nil {
                return nil, err
        }
        defer file.Close()
        var lines []string
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                lines = append(lines, scanner.Text())
        }
        return lines, scanner.Err()
}

func whichColors(item string, colorsTest[] string)([]string){
        var r[] string
        var itemSplit[] string = s.Split(item, ";")
        for _,element := range colorsTest {
                if s.Contains(element, itemSplit[3]) {
                        r = append(r, element)
                }
        }
        return r
}

func howMany(item string, colorsTest[] string)(int){
        var r int = 0
        var itemSplit[] string = s.Split(item, ";")
        for _,element := range colorsTest {
                if s.Contains(element, itemSplit[3]) {
                        fmt.Printf("   Checking this element : %f\n", element)
                        r++
                }
        }
        if r == 0 {
                r = 1
        }
        return r
}

func containsAny(item string, colorsTest[] string) (bool){
        var r = false;
        var itemSplit[] string = s.Split(item, ";")
        var colorsPost []string
        var itemSplit2 []string = s.Split(itemSplit[3], " ")
        for _,element := range colorsTest {
                for _, e := range itemSplit2 {
                        if s.ToLower(s.Replace(e, " ", "", -1)) != "" {
                                var trimmed[] string = s.SplitN(s.TrimLeft(s.Replace(element, ";", "; ", -1), ";"), ";", 2)
                                var shaved string = s.Replace(trimmed[0], " ", "", -1)
                                if shaved == e {
                                        colorsPost = s.Split(element, ";")
                                }
                        }
                }
        }
        for _,the_color := range colorsPost{
                if s.Replace(the_color, " ", "", -1) != "" {
                if s.Contains(the_color, "R ") != true{
                        if s.Contains(the_color, "G ") != true{
                                if s.Contains(the_color, "B ") != true{
                                                if s.Contains(the_color, "T ") != true{
                                                        var itempost = s.SplitN(s.Replace(item, ";", "; ", -1), ";", 4)
                                                        var coloritempost[] string = s.Split(itempost[3], " ")
                                                        for _,i := range coloritempost{
                                                                if s.Replace(i, " ", "", -1) != "" {
                                                                        if s.Replace(i, " ", "", -1) == s.Replace(the_color, " ", "", -1) {
                                                                                fmt.Printf("   Found existing color/alias : %f\n", the_color)
                                                                                r = true
                                                                        }else{
                                                                                fmt.Printf("   No known color/alias : %f\n", the_color)
                                                                        }
                                                                }
                                                        }
                                                }
                                        }
                                }
                        }
                }
        }
        return r
}

func generate (config string)(error){
        img := image.NewRGBA(image.Rect(0, 0, 32, 32))
        var file,err = readLines(config)
        var rv error
        var COLORS []string
        for _,element := range file {
                if s.Contains(element, "point") {
                        var cleanElement = s.Replace(element, "point", "", -1)
                        var splitElement = s.Split( cleanElement, ";" )
                        fmt.Printf(" Cleaned Element : %f\n", cleanElement)
                        var X int = -1
                        var Y int = -1
                        var R uint8 = 0
                        var G uint8 = 0
                        var B uint8 = 0
                        var T uint8 = 0
                        //fmt.Printf(" Preliminary Point Summary : %f\n", element)
                        for _, cleaned := range splitElement {
                                if s.Contains(cleaned, "X") {
                                        var pX,_ = strconv.ParseInt(s.Replace(cleaned, "X ", "", -1), 10, 8)
                                        X = int(pX)
                                        //fmt.Printf("  Point Summary X: %f\n", X)
                                }else if s.Contains(cleaned, "Y") {
                                        var pY,_ = strconv.ParseInt(s.Replace(cleaned, "Y ", "", -1), 10, 8)
                                        Y = int(pY)
                                        //fmt.Printf("  Point Summary Y: %f\n", Y)
                                }else if s.Contains(cleaned, "R") {
                                        var pR,_ = strconv.ParseInt(s.Replace(cleaned, "R ", "", -1), 10, 8)
                                        R = uint8(pR)
                                        //fmt.Printf("  Point Red : %f\n", R)
                                }else if s.Contains(cleaned, "G") {
                                        var pG,_ = strconv.ParseInt(s.Replace(cleaned, "G ", "", -1), 10, 8)
                                        G = uint8(pG)
                                        //fmt.Printf("  Point Green : %f\n", G)
                                }else if s.Contains(cleaned, "B") {
                                        var pB,_ = strconv.ParseInt(s.Replace(cleaned, "B ", "", -1), 10, 8)
                                        B = uint8(pB)
                                        //fmt.Printf("  Point Blue : %f\n", B)
                                }else if s.Contains(cleaned, "T") {
                                        var pT,_ = strconv.ParseInt(s.Replace(cleaned, "T ", "", -1), 10, 8)
                                        T = uint8(pT)
                                        //fmt.Printf("  Point Alpha : %f\n", T)
                                }else if containsAny(element, COLORS) {
                                        var count int = howMany(element, COLORS)
                                        var whichcolors[] string = whichColors(element, COLORS)
                                        fmt.Printf("-   Colors: %f\n", whichcolors)
                                        fmt.Printf("-   Count: %f\n", count)
                                        r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
                                        var times int = r.Intn(count) + 1
                                        for _, colorcleaned := range COLORS {
                                                var the_color = s.Split(colorcleaned, ";")                                                //if s.Contains(colorcleaned, splitElement[3]){
                                                        var pR,_ = strconv.ParseInt(s.Replace(the_color[2], "R ", "", -1), 10, 8)
                                                        R = uint8(pR)
                                                        //fmt.Printf("   Point Red(a) : %f\n", R)
                                                        var pG,_ = strconv.ParseInt(s.Replace(the_color[3], "G ", "", -1), 10, 8)
                                                        G = uint8(pG)
                                                        //fmt.Printf("   Point Green(a) : %f\n", G)
                                                        var pB,_ = strconv.ParseInt(s.Replace(the_color[4], "B ", "", -1), 10, 8)
                                                        B = uint8(pB)
                                                        //fmt.Printf("   Point Blue(a) : %f\n", B)
                                                        var pT,_ = strconv.ParseInt(s.Replace(the_color[5], "T ", "", -1), 10, 8)
                                                        T = uint8(pT)
                                                        //fmt.Printf("   Point Alpha(a) : %f\n", T)
                                                        fmt.Printf("-   Elapsed : %f\n", times)
                                                        times--
                                                        if times <= 0 {
                                                                break
                                                        }
                                        }
                                }
                        }
                        fmt.Printf("  Creating point : %f\n", cleanElement)
                        img.Set(X, Y, color.RGBA{R, G, B, T})
                }else if s.Contains(element, "color") {
                        element = s.Replace(element, "color", "", -1)
                        //fmt.Printf(" Found new color : %f\n", element)
                        COLORS = append(COLORS, element)
                }
        }
        // Save to file.png
        var name = s.Replace(s.Replace(s.Replace(config, "/", "", -1), ".txt", "", -1), "skel", "", -1) + ".png"
        fmt.Printf("Saving PNG file : %f\n", name)
        f, _ := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
        defer f.Close()
        png.Encode(f, img)
        if err != nil {  rv = err }
        return rv
}

func main() {
        var cfg = "config.txt"
        var file,_ = readLines(cfg)
        fmt.Printf("Using config file : %f\n", cfg)
        for _,element := range file {
                if s.Contains(element, "path=") {
                        var path=s.Replace(element, "path=", "", -1)
                        fmt.Printf(" Using skeleton file : %f\n", path)
                        generate(path)
                }
        }
}

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", count)
	appengine.Main()
}

func count(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("visits")
	v := 0
	if err != nil {
		c = &http.Cookie{
			Name:  "visits",
			Value: "1",
		}
		v = 1
		http.SetCookie(w, c)
	} else {
		v, _ = strconv.Atoi(c.Value)
		v++
		c.Value = strconv.Itoa(v)
		http.SetCookie(w, c)
	}
	fmt.Fprintln(w, "<p>Hello! You've been here "+c.Value+" times!</p></br>")

	talk(v, w)
}

func talk(v int, w http.ResponseWriter) {
	if v > 10 && v <= 20 {
		fmt.Fprintln(w, "You're persistent...")
	} else if v > 20 && v <= 30 {
		fmt.Fprintln(w, "Okay, yes, it's working. Move on.")
	} else if v > 30 && v <= 40 {
		fmt.Fprintln(w, "Don't you have anything better to do?")
	} else if v > 40 && v <= 50 {
		fmt.Fprintln(w, "Seriously, move on.")
	} else if v > 70 && v <= 80 {
		fmt.Fprintln(w, "Still here?")
	} else if v > 80 && v <= 90 {
		fmt.Fprintln(w, "What's wrong with you?")
	} else if v > 90 && v <= 100 {
		fmt.Fprintln(w, "Get a life, man.")
	} else if v > 100 && v <= 102 {
		fmt.Fprintln(w, "F_ it. I'm out. You stay if you want.")
	} else if v > 120 && v <= 130 {
		fmt.Fprintln(w, "Come on!!!")
	} else if v > 130 && v <= 140 {
		fmt.Fprintln(w, "Cut it out!")
	} else if v > 140 && v <= 150 {
		fmt.Fprintln(w, "I'm serious.")
	} else if v > 150 && v <= 160 {
		fmt.Fprintln(w, "You're annoying me.")
	} else if v > 160 && v <= 170 {
		fmt.Fprintln(w, "You don't have too many friends, do you?")
	} else if v > 170 && v <= 180 {
		fmt.Fprintln(w, "Okay, I see that you have a strong spirit. Thus I'll tell you a secret.")
	} else if v > 180 && v <= 190 {
		fmt.Fprintln(w, "It is a test.")
	} else if v > 190 && v <= 200 {
		fmt.Fprintln(w, "First 200 clicks filtered the weak. Let's see how far you will go.")
	} else if v > 210 && v <= 220 {
		fmt.Fprintln(w, "Good.")
	} else if v > 240 && v <= 250 {
		fmt.Fprintln(w, "Keep going.")
	} else if v > 250 && v <= 260 {
		fmt.Fprintln(w, "There's a deep meaning to all of this.")
	} else if v > 260 && v <= 270 {
		fmt.Fprintln(w, "And those weaklings will never know.")
	} else if v > 310 && v <= 320 {
		fmt.Fprintln(w, "Haha! Thought I was gone?")
	} else if v > 320 && v <= 330 {
		fmt.Fprintln(w, "We have only just begun, my friend.")
	} else if v > 330 && v <= 340 {
		fmt.Fprintln(w, "So how do you like that 'Game of Thrones'?")
	} else if v > 340 && v <= 350 {
		fmt.Fprintln(w, "Yeah, I know.")
	} else if v > 350 && v <= 360 {
		fmt.Fprintln(w, "Getting tired yet?")
	} else if v > 360 && v <= 370 {
		fmt.Fprintln(w, "Well, too bad. We have a long way ahead of us, my friend.")
	} else if v > 370 && v <= 380 {
		fmt.Fprintln(w, "So a drunk dog is trying to get back home through a locked door.")
	} else if v > 380 && v <= 390 {
		fmt.Fprintln(w, "Owner hears the noise and comes to the door.")
	} else if v > 390 && v <= 400 {
		fmt.Fprintln(w, "Drunk again?! You're useless animal!")
	} else if v > 400 && v <= 410 {
		fmt.Fprintln(w, "No dude! I'm totally sober!")
	} else if v > 400 && v <= 410 {
		fmt.Fprintln(w, "Sober? Yeah, right. You probable don't even remember how to bark!")
	} else if v > 410 && v <= 420 {
		fmt.Fprintln(w, "What?! Never!")
	} else if v > 420 && v <= 430 {
		fmt.Fprintln(w, "Yeah, then prove it! Bark!")
	} else if v > 430 && v <= 440 {
		fmt.Fprintln(w, "Sure! Bark-bark-bark-bark...")
	} else if v > 440 && v <= 450 {
		fmt.Fprintln(w, "Sorry...")
	} else if v > 450 && v <= 460 {
		fmt.Fprintln(w, "I really don't know why I like this joke...")
	} else if v > 460 && v <= 470 {
		fmt.Fprintln(w, "So anyway, how's life?")
	} else if v > 470 && v <= 480 {
		fmt.Fprintln(w, "Not that interesting from the looks of it...")
	} else if v > 480 && v <= 490 {
		fmt.Fprintln(w, "I mean you're still here.")
	} else if v > 490 {
		fmt.Fprintln(w, "To be continued...")
	}
}

package lovelettercreator

import(
    _"fmt"
    "math/rand"
    "strings"
    "time"
)

func CreateLetter(senderGender bool, senderName string, recieveName string) string{
    // senderGender = 1 : female
    // senderGender = 0 : male
    rand.Seed(time.Now().UnixNano())

    var getRand = func(a []string)string{
        return a[rand.Intn(len(a))]
    }

    var salutation string = "My "+strings.Title(getRand(recipAdj()))+" "
    var line1 string= "My " + getRand(senderAdj()) + " " + getRand(senderNouns()) + " " +getRand(adverbs())+" "+getRand(verbs()) + " your " + getRand(recipNouns()) + "."
    var line2 string= "You are my " + getRand(senderAdj()) +" " + getRand(senderNouns())+"."
    var line3 string= "My " + getRand(senderNouns()) + " " + getRand(verbs()) + " for your " + getRand(recipNouns())+"."
    var line4 string= "Your " + getRand(bodyParts(!senderGender))+" is my "+ getRand(senderAdj())+" "+ getRand(senderNouns()) + "."
    var line5 string= "I want to wake up to your "+ getRand(recipAdj()) + " " + getRand(bodyParts(!senderGender)) + ", "+ getRand(recipAdj()) + " " + getRand(bodyParts(!senderGender)) + ",and "+ getRand(recipAdj()) + " " + getRand(bodyParts(!senderGender)) + "."
    var line6 string= "You are my "+ getRand(senderAdj()) +" " + getRand(senderNouns()) + ": A " +getRand(recipAdj()) + " " +getRand(recipNouns()) +"."
    if senderName == ""{
            senderName += "Your " + strings.Title(getRand(senderAdj())) +" "+strings.Title(getRand(senderNouns()))+"."
    }
    var closing string = getRand(closeWords()) + ",\n"+ senderName

    if recieveName == ""{
        if senderGender {
            salutation += strings.Title(getRand(altSalMale())) //female is writing 
        }else{
            salutation += strings.Title(getRand(altSalFemale())) //male is writing
        }
    }else{
        salutation += strings.Title(recieveName)
    }
    salutation += ","

    bodyList := []string{
        line1,
        line2,
        line3,
        line4,
        line5,
        //say(senderGender),
        func (g bool) string {
            if g {
                return getRand(sheSays())
            }else {
                return getRand(heSays())
            }
        }(senderGender),
        getRand(genSetOne()),
        getRand(genSetTwo()),
        line6,
    }

    letterBody := strings.Join(bodyList," ")

    letterList := []string {
        salutation,
        letterBody,
        closing,
    }

    letter := strings.Join(letterList, "\n")
    return letter
}

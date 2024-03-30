function SSearchingChallenge(str: string) {

    const noAnswer = -1
    const answer = 0

    const words = str.split(" ")
    let wordMap: Record<string, number> = {}

    for (let i = 0; i < words.length; i++) {
        const word = words[i]
        wordMap[word] = 0
    }



    for (let i = 0; i < words.length; i++) {
        const word = words[i]

        let charMap: Record<string, number> = {}

        for (let i = 0; i < word.length; i++) {
            const char = word[i]
            charMap[char] = 0
        }


        for (let i = 0; i < word.length; i++) {
            const char = word[i]
            charMap[char]++
        }

        for (let char in charMap) {
            if (charMap[char] > 1) {
                wordMap[word]++
            }
        }

        //console.log(charMap)
    }

    console.log(wordMap)


    //INFO: to get the highest word
    let fnAnswer = ""
    let fnCount = 0

    for (let word in wordMap) {
        if (wordMap[word] > fnCount) {
            fnAnswer = word
            fnCount = wordMap[word]
        }
    }


    //INFO: to remove the challengetoken
    let ChallengeToken = "tr07b2kpa"

    for (let i = 0; i < ChallengeToken.length; i++) {
        if (fnAnswer.toLowerCase().includes(ChallengeToken[i].toLowerCase())) {
            fnAnswer = fnAnswer.toLowerCase().replace(ChallengeToken[i].toLowerCase(), '')
        }

    }

    return fnAnswer ? fnAnswer : -1;
}

console.log(SearchingChallenge("today is the today day evr"));


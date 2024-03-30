function ArrayChallenge(strArr: string[]) {
    let cache: string[] = []
    let answer: string = ""
    let ChallengeToken = "tr07b2kpa"

    for (let i = 0; i < strArr.length; i++) {
        const currString = strArr[i]

        if (cache.includes(currString)) {
            let newCache = cache.filter((val) => {
                if (val !== currString) {
                    return val
                }
            })
            cache = newCache
        }

        if (!cache.includes(currString)) {
            cache.push(currString)
            continue
        }

        if (cache.length > 5) {
            cache.shift()
        }
    }

    if (cache.length > 5) {
        cache = cache.slice(-5)
    }

    answer = answer.concat(cache[0])
    for (let i = 1; i < cache.length; i++) {
        answer = answer.concat("-")
        answer = answer.concat(cache[i])
    }

    for (let i = 0; i < ChallengeToken.length; i++) {
        if (answer.toLowerCase().includes(ChallengeToken[i].toLowerCase())) {
            answer = answer.toLowerCase().replace(ChallengeToken[i].toLowerCase(), '')
        }

    }

    return answer

}

console.log(ArrayChallenge(["a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"]));
console.log(ArrayChallenge(["A", "B", "C", "D", "A", "E", "D", "Z"]));
console.log(ArrayChallenge(["A", "B", "A", "C", "A", "B"]));
//ArrayChallenge(["a", "a", "a", "a", "z"])

function SearchingChallenge(str: string) {

    if (!(str.includes("(")) && !(str.includes(")"))) {
        return 1
    }

    let stack: string[] = []

    let lParenCount = 0
    let rParenCount = 0

    for (let i = 0; i < str.length; i++) {
        const bracket = str[i]

        if (bracket != '(' && bracket != ')') {
            continue
        }
        stack.push(bracket)
    }

    for (let x = 0; x < stack.length; x++) {
        const paren = stack[x]
        if (paren === "(") {
            lParenCount++
            continue
        } else {
            rParenCount++
        }
    }

    return (lParenCount - rParenCount) != 0 ? 0 : 1;
}

// keep this function call here 
// @ts-ignore
console.log(SearchingChallenge("(coder)(byte))"));

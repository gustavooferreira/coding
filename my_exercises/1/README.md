# Gus 1. Match parenthesis

Given a string (array of characters), return whether there are matching parenthesis in the string.

The parenthesis in the string can be of all 3 types: `()`, `[]`, `{}`.

Some or all types of parenthesis may show up in a string.

**Example 1:**

> Input: str = "(this is a phrase {that contains another phrase [123]})"
>
> Output: true
>
> Explanation: Every single parenthesis that is opened, is met with the parenthesis close in the right order.

**Example 2:**

> Input: str = "this [ is not match by the right parenthesis close)"
>
> Output: false

**Example 3:**

> Input: str = "([{[]}])[]{{()}}"
>
> Output: true

**Follow-up:** Keep track of how many matching brackets of each kind show up and print the result.

**Example:**

> Input: str = "([{[]}])[]{{()}}"
>
> Output: true
>
> Stdout:
> () -> 2
> [] -> 3
> {} -> 3

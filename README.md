[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/LuIogM4K)
#### User Stories and Acceptance Criteria

#### ======================================

User Story: As an user, I want to know the name and the purpose of the spaced repetition application, so that I can decide whether I want to continue interacting with it.

    Acceptance Criteria:
        When I execute the spaced repetition CLI, I see the name of the learning app, and a  brief description
        of what the app does.

---

User Story: As an user, I want to create a record so that I can use the app.

    Acceptance Criteria:
        When I visit the the spaced repetition CLI for the first time, I should be directed to a registration prompt
        On that page, I can enter my name it should be at least of 5 characters long
        If the information is correct, upon submitting the info, I should be taken to the login screen
        If the information is incorrect, I should be given proper error messages and a prompt to enter the correct information

---

User Story: As a user, I should be able to login, so that I can use the app

    Acceptance criteria:
        When I execute the spaced repetition CLI, I should be able to skip the registration and login using my record name
        If my name exists then the app should greet me with a message
        "Hello MY_NAME"
        If my name is incorrect then the app should reject my information and give me a meaningful message so I can try again or go back to register

---

User Story: As a user, I am presented with a word so that I can learn the meaning of the word

    Acceptance criteria:
        When I log into my spaced repetition app, I am presented with a word that I want to learn
        I can type the meaning of the word
        When I press enter, I submit my answer

---

User Story: As a user, I am given feedback to my answer so that I learn the word

    Acceptance criteria:
        When I submit the answer, I should get a response for correct answer
        If the answer is wrong, I should get a feedback as well as the correct answer so I can learn the correct answer and get it right the next time.

---

User Story: As a user, when I am using the app, I should be able to move to the next word so I can learn more words

    Acceptance criteria:
        When I learn a word, I should see a next option to move to the next word

---

User Story: As a user, when I am using the app, my progress should be recorded so I know which questions I got correct or wrong

    Acceptance criteria:
        When I get a word correct, I should get a score if I get that word correct
        When I get the word wrong, my score won't change

---

User Story: As a user, I should see the words based on the spaced repetition algorithm, so that I can learn by frequently seeing the word that I have hard time remembering

    Acceptance criteria:
        When I get a word correct, I should see that word at a later time
        When I get the word wrong, I should see the word more frequently

---

User Story: As a user, I should be able to move through the words for as long as I want so that I can learn for as long as I want

    Acceptance criteria:
        Once I log in and start learning, I should be able to see the words based on the spaced repetition algorithm
        I can choose to finish or reset at any time
        Otherwise, unless I choose finish, I should be able to move through the words without any ending

---

User Story: As a user, I should be able to see my overall progress so that I know which words I need to learn more

    Acceptance criteria:
        When I recover my account, I should be able to see a progress option and see my overall progress ordered by retention rate
        As I am learning, I should be able to choose the progress option and see my overall progress updated reflecting my current progress

---

User Story: As a user, I should be able to successfully finish of the spaced repetition app so that I can rest my brain

    Acceptance criteria:
        When I choose the finish option, all my session data should be removed from the screen and I should
        see a message that I have successfully finished a session

## Example screens

1.

```
This is Spaced repetition app

It helps you to learn english
using the spaced repetition
algorithm!!!

Good luck!
```

2.

```
I don't know who you are...
Did you want to:
1. Register
2. Login
```

3. If register

```
Please tell me your name?
>
This is not a valid name
Please tell me your name?
> Wences
Welcome to the Spaced Repetition App!
```

3. If login

```
Please remember me your name?
> Wen
Invalid user name
Please remember me your name?
> Wences
Hello Wences!
```

4. Learn

```
What does learning means in Spanish?
> aprender

Great! What did you want to do?
1. Go to the next word
2. See my progress
3. Finish
```

4. Progress

```
Your current progress is:
Words       | Retention Rate
learning    | 100%
walking     | 100%
certain     | 75%
intricate   | 50%
calculation | 0%
...

Options:
1. Go to the next word
2. See my progress
3. Finish
```

4. Finish

```
Thanks for practicing with the
Space repetition app
See you later!
```

## Spaced repetition example

[spaced.md](spaced.md)

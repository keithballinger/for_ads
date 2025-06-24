# Architecture Design: A Love Letter Website

This document outlines the technical architecture for a website that presents a simple puzzle and, upon solving, reveals a Shakespearean sonnet.

## 1. System Overview

The system is composed of two main parts:

*   **Frontend:** A Svelte single-page application (SPA) responsible for the user interface, puzzle logic, and communication with the backend.
*   **Backend:** A Go server that provides the sonnet via a simple API.

## 2. Frontend Architecture

*   **Framework:** Svelte will be used for its simplicity and performance.
*   **Styling:** Mantine will be used for its rich set of components and aesthetically pleasing default styles. The application will have a romantic and elegant theme.
*   **Structure:**
    *   `main.js`: The entry point of the Svelte application.
    *   `App.svelte`: The main component that will handle routing between the puzzle and the sonnet pages.
    *   `Puzzle.svelte`: A component for the initial puzzle page. The puzzle will be a simple question-and-answer.
    *   `Sonnet.svelte`: A component to display the sonnet. This page will be beautifully styled with a focus on typography and layout.

## 3. Backend Architecture

*   **Language:** Go will be used for its performance and simplicity.
*   **Framework:** The standard library's `net/http` package will be sufficient for this simple API.
*   **API Endpoints:**
    *   `GET /api/sonnet`: This endpoint will return a JSON object containing the Shakespearean sonnet.
        *   **Response:**
            ```json
            {
              "title": "Sonnet 116",
              "lines": [
                "Let me not to the marriage of true minds",
                "Admit impediments. Love is not love",
                "Which alters when it alteration finds,",
                "Or bends with the remover to remove:",
                "O no! it is an ever-fixed mark",
                "That looks on tempests and is never shaken;",
                "It is the star to every wandering bark,",
                "Whose worth's unknown, although his height be taken.",
                "Love's not Time's fool, though rosy lips and cheeks",
                "Within his bending sickle's compass come:",
                "Love alters not with his brief hours and weeks,",
                "But bears it out even to the edge of doom.",
                "If this be error and upon me proved,",
                "I never writ, nor no man ever loved."
              ]
            }
            ```

## 4. Project Structure

```
.
├── backend
│   └── main.go
├── frontend
│   ├── public
│   │   └── index.html
│   ├── src
│   │   ├── App.svelte
│   │   ├── Puzzle.svelte
│   │   ├── Sonnet.svelte
│   │   └── main.js
│   └── package.json
├── arch.md
└── tasks.md
```

import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import { createGlobalStyle } from "styled-components";
import { ThemeProvider } from "@mui/material/styles";
import theme from "@/theme/Theme.ts";
import CssBaseline from "@mui/material/CssBaseline";
import { BrowserRouter as Router } from "react-router-dom";
import { AuthProvider } from "@/context/AuthContext";
import { APIProvider } from "./context/APIContext";

// define global styles
const GlobalStyle = createGlobalStyle`
  body {
    margin: 0;
    padding: 0;
    font-family: sans-serif;
  }
`;

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <GlobalStyle />
      <CssBaseline />
      <Router>
        <APIProvider>
          <AuthProvider>
            <App />
          </AuthProvider>
        </APIProvider>
      </Router>
    </ThemeProvider>
  </React.StrictMode>
);

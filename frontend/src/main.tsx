import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import { ThemeProvider } from "@mui/material/styles";
import theme from "@/theme/Theme.ts";
import CssBaseline from "@mui/material/CssBaseline";
import { BrowserRouter as Router } from "react-router-dom";
import { APIProvider } from "@/context/APIContext";
import AuthProvider from "@/context/AuthContext";
import GlobalStyle from "@/GlobalStyle";
import { SnackbarProvider } from "./context/SnackbarContext";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <GlobalStyle />
      <CssBaseline />
      <SnackbarProvider>
        <Router>
          <APIProvider>
            <AuthProvider>
              <App />
            </AuthProvider>
          </APIProvider>
        </Router>
      </SnackbarProvider>
    </ThemeProvider>
  </React.StrictMode>
);

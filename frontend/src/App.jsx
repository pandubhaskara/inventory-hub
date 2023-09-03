import { BrowserRouter } from "react-router-dom";
import { HelmetProvider } from "react-helmet-async";
// component
import { StyledChart } from "./components/chart";
import ScrollToTop from "./components/scroll-to-top";
// router
import Router from "./Route.jsx";
// theme
import "./App.css";
import ThemeProvider from "./theme/index.jsx";

function App() {
  return (
    <HelmetProvider>
      <BrowserRouter>
        <ThemeProvider>
          <ScrollToTop />
          <StyledChart />
          <Router />
        </ThemeProvider>
      </BrowserRouter>
    </HelmetProvider>
  );
}

export default App;

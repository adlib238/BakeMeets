import "./App.css";
import Home from "./pages/Home";
import NoMatch from "./pages/Nomatch";
import Favorite from "./pages/Favorite";
import Profile from "./pages/Profile";
import AppLayout from "./components/layout/AppLayout";

import { Route, BrowserRouter as Router, Routes } from "react-router-dom";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<AppLayout />}>
          <Route index element={<Home />} />
          <Route path="/favorite" element={<Favorite />} />
          <Route path="/profile" element={<Profile />} />
          <Route path="*" element={<NoMatch />} />
        </Route>
      </Routes>
    </Router>
  );
}

export default App;

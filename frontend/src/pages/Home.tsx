import React from "react";
import { Box } from "@mui/material";
import BreadGallery from "../components/BreadGallery";

const Home = () => {
  return (
    <Box sx={{ display: "flex" }}>
      <Box sx={{ flexGrow: 1 }}>
        <BreadGallery />
      </Box>
    </Box>
  );
};

export default Home;

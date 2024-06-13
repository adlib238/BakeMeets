import React, { useState } from "react";
import breadData from "../data/breadData";
import { Bread } from "../types/Bread";
import { Box, Typography, IconButton, Modal } from "@mui/material";
import FavoriteBorderIcon from "@mui/icons-material/FavoriteBorder";
import CommentIcon from "@mui/icons-material/Comment";

const styles = {
  breadGallery: {
    display: "flex",
    flexWrap: "wrap",
    justifyContent: "space-between",
    gap: 2,
  },
  breadCard: {
    flex: "1 1 calc(33.333% - 16px)",
    boxSizing: "border-box",
    border: "1px solid #ccc",
    borderRadius: 1,
    overflow: "hidden",
    mb: 2,
    cursor: "pointer",
  },
  breadImage: {
    width: "100%",
    height: "auto",
  },
  breadInfo: {
    p: 2,
  },
  breadStats: {
    display: "flex",
    justifyContent: "space-between",
  },
  modalStyle: {
    position: "absolute",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    width: 400,
    bgcolor: "background.paper",
    border: "2px solid #000",
    boxShadow: 24,
    p: 4,
  },
};

const BreadDetailModal = ({ bread, open, onClose }) => (
  <Modal open={open} onClose={onClose}>
    <Box sx={styles.modalStyle}>
      <Typography variant="h6">{bread.name}</Typography>
      <img src={bread.imageUrl} alt={bread.name} style={styles.breadImage} />
      <Typography>{bread.description}</Typography>
      <Box sx={styles.breadStats}>
        <Box>
          <IconButton>
            <FavoriteBorderIcon />
          </IconButton>
          {bread.likes}
        </Box>
        <Box>
          <IconButton>
            <CommentIcon />
          </IconButton>
          {bread.comments}
        </Box>
      </Box>
    </Box>
  </Modal>
);

const BreadGallery = () => {
  const [selectedBread, setSelectedBread] = useState(null);
  const [modalOpen, setModalOpen] = useState(false);

  const handleCardClick = (bread) => {
    setSelectedBread(bread);
    setModalOpen(true);
  };

  const handleModalClose = () => {
    setModalOpen(false);
    setSelectedBread(null);
  };

  return (
    <Box sx={{ display: "flex" }}>
      <Box sx={styles.breadGallery}>
        {breadData.map((bread: Bread) => (
          <Box
            key={bread.id}
            sx={styles.breadCard}
            onClick={() => handleCardClick(bread)}
          >
            <img
              src={bread.imageUrl}
              alt={bread.name}
              style={styles.breadImage}
            />
            <Box sx={styles.breadInfo}>
              <Typography>{bread.name}</Typography>
              <Box sx={styles.breadStats}>
                <Box>
                  <FavoriteBorderIcon /> {bread.likes}
                </Box>
                <Box>
                  <CommentIcon /> {bread.comments}
                </Box>
              </Box>
            </Box>
          </Box>
        ))}
      </Box>
      {selectedBread && (
        <BreadDetailModal
          bread={selectedBread}
          open={modalOpen}
          onClose={handleModalClose}
        />
      )}
    </Box>
  );
};

export default BreadGallery;

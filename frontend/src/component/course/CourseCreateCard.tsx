import { AddCircleOutline } from "@mui/icons-material";
import { Box, IconButton, Typography } from "@mui/material";
import React from "react";

const customStyle = {
  createCourseDiv: {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    margin: "8px",
    borderRadius: "24px",
    borderStyle: "dashed",
    borderWidth: "2px",
    width: "100%",
    maxWidth: "280px",
    minWidth: "260px",
    height: "240px",
    color: "#6B6F7B",
  },
  createCourseButton: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
  },
  createCourseColor: {
    color: "#686F7B",
  },
  createCourseTitle: {
    fontSize: "14px",
    fontWeight: "bold",
    fontFamily: "Poppins",
    marginLeft: "8px",
  },
};

export interface ICourseCreateCardProps {}

const CourseCreateCard = () => {
  return (
    <div style={customStyle.createCourseDiv}>
      <Box sx={customStyle.createCourseButton}>
        <IconButton size="large">
          <AddCircleOutline
            sx={customStyle.createCourseColor}
            fontSize="large"
          />
        </IconButton>
        <Typography
          noWrap={true}
          variant="subtitle1"
          sx={{
            ...customStyle.createCourseColor,
            ...customStyle.createCourseTitle,
          }}
        >
          Add new class
        </Typography>
      </Box>
    </div>
  );
};

export default CourseCreateCard;

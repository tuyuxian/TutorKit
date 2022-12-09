import { MoreVert, Paid, Person, WatchLater } from "@mui/icons-material";
import {
  Avatar,
  AvatarGroup,
  Box,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  IconButton,
  Typography,
} from "@mui/material";
import React from "react";

const customStyle = {
  courseCardDiv: {
    display: "inline-block",
    justifyContent: "space-around",
    alignContent: "center",
    margin: "8px",
    boxShadow: "0px 0px 8px rgba(0,0,0,0.2)",
  },
  card: {
    width: "280px",
    height: "240px",
    borderRadius: "24px",
  },
  avatar: {
    width: 36,
    height: 36,
    fontSize: 14,
  },
  cardHeaderColor: {
    color: "#586D80",
  },
  cardHeaderContent: {
    display: "flex",
    flexDirection: "column",
    alignItems: "flex-start",
  },
  cardHeaderTitle: {
    fontSize: "18px",
    fontWeight: "bold",
    textOverflow: "ellipsis",
    overflow: "hidden",
    maxWidth: "220px",
  },
  cardHeaderSubHeader: {
    fontSize: "14px",
  },
  cardContentBox: {
    display: "flex",
    flexDirection: "column",
    justifyContent: "space-between",
    height: "163px",
  },
  cardContentRow: {
    display: "flex",
    alignItems: "center",
    margin: "8px 0 8px 0",
  },
  cardContentColor: {
    color: "#9EAEC5",
  },
  cardContentTitle: {
    fontSize: "14px",
    fontWeight: "bold",
    marginLeft: "8px",
  },
};

export interface ICourseCardProps {
  /**
   * tutor, student, and parents.
   */
  usage: string;
}

const CourseCard: React.FC<ICourseCardProps> = ({ usage }) => {
  return (
    <div style={{ ...customStyle.card, ...customStyle.courseCardDiv }}>
      <Card sx={{ ...customStyle.card, boxShadow: "none" }}>
        <CardHeader
          action={
            <IconButton aria-label="settings">
              <MoreVert />
            </IconButton>
          }
          title="Calculus"
          titleTypographyProps={{
            ...customStyle.cardHeaderColor,
            ...customStyle.cardHeaderTitle,
          }}
          subheader="08/31/2022 - 10/31/2022"
          subheaderTypographyProps={{
            ...customStyle.cardHeaderSubHeader,
            ...customStyle.cardHeaderColor,
          }}
          sx={{
            "& .MuiCardHeader-content": {
              ...customStyle.cardHeaderContent,
            },
          }}
        />
        <Box sx={customStyle.cardContentBox}>
          <CardContent>
            <Box sx={customStyle.cardContentRow}>
              <WatchLater sx={customStyle.cardContentColor} />
              <Typography
                noWrap={true}
                variant="subtitle1"
                sx={{
                  ...customStyle.cardContentColor,
                  ...customStyle.cardContentTitle,
                }}
              >
                Tue, Oct 4 16: 30
              </Typography>
            </Box>
            {usage === "tutor" || usage === "parents" ? (
              <Box sx={customStyle.cardContentRow}>
                <Paid sx={customStyle.cardContentColor} />
                <Typography
                  noWrap={true}
                  variant="subtitle1"
                  sx={{
                    ...customStyle.cardContentColor,
                    ...customStyle.cardContentTitle,
                  }}
                >
                  $600 /hrs
                </Typography>
              </Box>
            ) : null}
            {usage === "parents" || usage === "student" ? (
              <Box sx={customStyle.cardContentRow}>
                <Person sx={customStyle.cardContentColor} />
                <Typography
                  noWrap={true}
                  variant="subtitle1"
                  sx={{
                    ...customStyle.cardContentColor,
                    ...customStyle.cardContentTitle,
                  }}
                >
                  Bruno Mars
                </Typography>
              </Box>
            ) : null}
          </CardContent>
          {usage === "tutor" && (
            <CardActions disableSpacing>
              <AvatarGroup
                max={4}
                sx={{
                  "& .MuiAvatar-root": { ...customStyle.avatar },
                }}
              >
                <Avatar alt="Remy Sharp" src="/static/images/avatar/1.jpg" />
                <Avatar alt="Travis Howard" src="/static/images/avatar/2.jpg" />
                <Avatar alt="Cindy Baker" src="/static/images/avatar/3.jpg" />
                <Avatar alt="Agnes Walker" src="/static/images/avatar/4.jpg" />
                <Avatar
                  alt="Trevor Henderson"
                  src="/static/images/avatar/5.jpg"
                />
              </AvatarGroup>
            </CardActions>
          )}
        </Box>
      </Card>
    </div>
  );
};

export default CourseCard;

import { Helmet } from "react-helmet-async";
// @mui
import { useTheme } from "@mui/material/styles";
import { Grid, Container, Typography } from "@mui/material";
// sections
import {
  AppCurrentVisits,
  AppWidgetSummary,
} from "../sections/@dashboard/app";

// ----------------------------------------------------------------------

export default function DashboardAppPage() {
  const theme = useTheme();
  const account = {
    username: "pandu",
    role: "Admin",
    email: "pandu@demo.com",
    photoURL:
      "https://upload.wikimedia.org/wikipedia/commons/4/44/210604_%EA%B3%A0%EC%9C%A4%EC%A0%95%282%29.jpg",
  };

  return (
    <>
      <Helmet>
        <title> Dashboard | InventoryHub </title>
      </Helmet>

      <Container maxWidth="xl">
        <Typography variant="h4" sx={{ mb: 5 }}>
          Hi, {account.username}
        </Typography>

        <Grid container spacing={3}>
          <Grid item xs={12} sm={6} md={6}>
            <AppWidgetSummary
              title="In"
              total={20}
              color="info"
              icon={"ant-design:caret-down-filled"}
            />
          </Grid>

          <Grid item xs={12} sm={6} md={6}>
            <AppWidgetSummary
              title="Out"
              total={10}
              color="error"
              icon={"ant-design:caret-up-filled"}
            />
          </Grid>

          <Grid item xs={12} md={12} lg={12}>
            <AppCurrentVisits
              chartData={[
                { label: "In", value: 20 },
                { label: "Out", value: 10 },
              ]}
              chartColors={[
                theme.palette.info.main,
                theme.palette.error.main,
              ]}
            />
          </Grid>
        </Grid>
      </Container>
    </>
  );
}

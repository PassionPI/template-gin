import LocationSub from "@/components/LocationSub";
import Navigator from "@/components/Navigator";
import { getToken, invokePubKey } from "@/services/login";
import { TreeToRoute } from "@/utils/staticRoute";
import { suspense } from "@/utils/suspense";
import { Outlet, createBrowserRouter, redirect } from "react-router-dom";

export const ROUTE = TreeToRoute({
  login: null,
  sign: null,
  home: null,
});
export const HOME_ROUTE = ROUTE.home.__;

const Login = suspense(() => import("@/pages/Login"));
const Sign = suspense(() => import("@/pages/Sign"));
const Home = suspense(() => import("@/pages/Home"));
const NotFound = suspense(() => import("@/components/NotFound"));

export const AppRouter = createBrowserRouter([
  {
    path: "/",
    element: (
      <>
        <LocationSub />
        <Navigator />
        <Outlet />
      </>
    ),
    children: [
      {
        path: "/",
        loader: async () => {
          invokePubKey();
          if (getToken()) {
            return redirect(ROUTE.home.__);
          }
          return null;
        },
        children: [
          {
            path: ROUTE.login.__,
            element: <Login />,
          },
          {
            path: ROUTE.sign.__,
            element: <Sign />,
          },
        ],
      },
      {
        path: "/",
        children: [
          {
            index: true,
            loader: async () => {
              const token = getToken();
              if (!token) {
                return redirect(ROUTE.login.__);
              }
              return redirect(ROUTE.home.__);
            },
          },
          {
            path: ROUTE.home.__,
            children: [
              {
                index: true,
                element: <Home />,
              },
            ],
          },
        ],
      },
    ],
  },

  {
    path: "*",
    element: <NotFound />,
  },
]);

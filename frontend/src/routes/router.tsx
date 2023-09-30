import Navigator from "@/components/Navigator";
import { invokePubKey, isAuthorization } from "@/services/login";
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
        <Navigator />
        <Outlet />
      </>
    ),
    loader: async () => {
      invokePubKey();
      const { pathname } = location;
      const authorization = await isAuthorization();
      await new Promise((resolve) => setTimeout(resolve, 1000));
      console.log("authorization", authorization);
      if (authorization) {
        return null;
      }
      if (pathname === ROUTE.login.__ || pathname === ROUTE.sign.__) {
        return null;
      }
      return redirect(ROUTE.login.__);
    },
    children: [
      {
        path: "/",
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

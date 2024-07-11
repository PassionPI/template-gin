import ErrorBoundary from "@/components/ErrorBoundary";
import { AppRouter } from "@/routes/router";
import { ConfigProvider } from "antd";
import { extend } from "dayjs";
import duration from "dayjs/plugin/duration";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { Provider as ReduxProvider } from "react-redux";
import { RouterProvider } from "react-router-dom";
import { store } from "./app";
import "./index.css";

extend(duration);

const container = document.getElementById("root");

if (container) {
  const root = createRoot(container);
  root.render(
    <StrictMode>
      <ErrorBoundary>
        <ReduxProvider store={store}>
          <ConfigProvider theme={{ token: { colorPrimary: "#0367c4" } }}>
            <RouterProvider router={AppRouter} />
          </ConfigProvider>
        </ReduxProvider>
      </ErrorBoundary>
    </StrictMode>,
  );
} else {
  window.console.error("[id=root] element not found");
}

import { ROUTE } from "@/routes/router";
import type { FC } from "react";
import { memo } from "react";
import { useNavigate } from "react-router-dom";

const Home: FC = memo(() => {
  const nav = useNavigate();

  return (
    <div>
      {ROUTE.home.__}
      <button onClick={() => nav(ROUTE.login.__)}>/login</button>
    </div>
  );
});

Home.defaultProps = {};

export default Home;

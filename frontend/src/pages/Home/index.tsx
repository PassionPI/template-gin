import { ROUTE } from "@/routes/router";
import { request } from "@/utils/request";
import type { FC } from "react";
import { memo, useEffect } from "react";
import { useNavigate } from "react-router-dom";

const Home: FC = memo(() => {
  const nav = useNavigate();
  useEffect(() => {
    request({
      method: "post",
      url: "/api/ping",
    })
      .then(([err, res, meta]) => {
        if (err != null) {
          return err.message;
        } else {
          return res;
        }
      })
      .then((msg) => {
        console.log(msg);
      });
  }, []);
  return (
    <div>
      {ROUTE.home.__}
      <button onClick={() => nav(ROUTE.login.__)}>/login</button>
    </div>
  );
});

Home.defaultProps = {};

export default Home;

import { HOME_ROUTE } from "@/routes/router";
import { Button, Result } from "antd";
import type { FC } from "react";
import { memo } from "react";
import { useNavigate } from "react-router-dom";

const NotPassAuth: FC = memo(() => {
  const nav = useNavigate();
  return (
    <Result
      status="404"
      title="404"
      subTitle="Sorry, this page does not exist."
      extra={
        <Button type="primary" onClick={() => nav(HOME_ROUTE)}>
          Back Home
        </Button>
      }
    />
  );
});

NotPassAuth.defaultProps = {};

export default NotPassAuth;

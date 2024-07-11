import { useAppSelector } from "@/app";
import { FC, memo, useEffect } from "react";
import { useNavigate } from "react-router-dom";

const Navigator: FC = memo(() => {
  const nav = useNavigate();
  const { to, replace } = useAppSelector((state) => state.history);

  useEffect(() => {
    if (to) {
      nav(to, { replace });
    }
  }, [nav, to, replace]);

  return null;
});

Navigator.defaultProps = {};

export default Navigator;

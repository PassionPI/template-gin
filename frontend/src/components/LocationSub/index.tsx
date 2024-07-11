import type { FC } from "react";
import { memo, useEffect } from "react";
import { useLocation } from "react-router-dom";

type LocationSubscriber = (location: ReturnType<typeof useLocation>) => void;

const locationSubscribers: LocationSubscriber[] = [];

export const subscribeLocation = (...fns: LocationSubscriber[]) => {
  locationSubscribers.push(...fns);
};

const LocationSub: FC = memo(() => {
  const location = useLocation();

  useEffect(() => {
    locationSubscribers.forEach((fn) => fn(location));
  }, [location]);

  return null;
});

LocationSub.defaultProps = {};

export default LocationSub;

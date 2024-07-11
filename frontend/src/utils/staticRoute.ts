const __ = "__";

type BaseValue = { [__]: string };
type BaseTree = {
  [path: string]: null | BaseTree;
};
type Tree<T extends BaseTree> = {
  [K in keyof T]: T[K] extends null
    ? BaseValue
    : T[K] extends BaseTree
    ? Tree<T[K]> & BaseValue
    : never;
};

export function TreeToRoute<T extends BaseTree>(tree: T) {
  function parseRoute<T extends BaseTree>(route: T, base = ""): Tree<T> {
    return new Proxy(route, {
      get(target, p, receiver) {
        const next = Reflect.get(target, p, receiver);
        const path = `${base}/${p as string}`;

        if (p === __) {
          return base;
        }
        if (next === null) {
          return {
            [__]: path,
          };
        }
        if (next === undefined) {
          return;
        }

        return parseRoute(next as BaseTree, path);
      },
    }) as unknown as Tree<T>;
  }

  return parseRoute(tree, "");
}

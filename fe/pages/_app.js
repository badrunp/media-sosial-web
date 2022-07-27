import { SessionProvider, useSession } from "next-auth/react";
import { useRouter } from "next/router";
import { createContext, useContext, useEffect, useReducer } from "react";
import "../styles/globals.css";

const initialState = {
  user: undefined,
};
export const Store = createContext(initialState);
export function useStore() {
  return useContext(Store);
}

export const types = {
  SET_USER: "SET_USER",
};

function reducer(state, action) {
  switch (action.type) {
    case types.SET_USER:
      return {
        ...state,
        user: action.payload.user,
      };
    default:
      return state;
  }
}

function StoreProvider({ children }) {
  const [state, dispatch] = useReducer(reducer, initialState);
  return (
    <Store.Provider value={{ state, dispatch }}>{children}</Store.Provider>
  );
}

export default function App({
  Component,
  pageProps: { session, ...pageProps },
}) {
  return (
    <StoreProvider>
      <SessionProvider session={session}>
        <Render Component={Component} pageProps={pageProps} />
      </SessionProvider>
    </StoreProvider>
  );
}

function Render({ Component, pageProps }) {
  return (
    <>
      {Component.auth ? (
        <Auth>
          <Component {...pageProps} />
        </Auth>
      ) : Component.guest ? (
        <Guest>
          <Component {...pageProps} />
        </Guest>
      ) : (
        <Component {...pageProps} />
      )}
    </>
  );
}

function Auth({ children }) {
  const router = useRouter();
  const { data: session, status } = useSession();
  const { dispatch } = useStore();

  useEffect(() => {
    if (session) {
      dispatch({ type: types.SET_USER, payload: { user: session.user } });
    }
  }, [session]);

  if (status == "loading") {
    return <p>loading.. auth</p>;
  }

  if (status != "loading") {
    if (status == "unauthenticated") {
      router.push("/signin");
      return;
    } else if (status == "authenticated") {
      return children;
    }
  }
}

function Guest({ children }) {
  const router = useRouter();
  const { status } = useSession();
  const { dispatch } = useStore();

  useEffect(() => {
    dispatch({ type: types.SET_USER, payload: { user: null } });
  }, []);

  if (status == "loading") {
    return <p>loading.. guest</p>;
  }

  if (status != "loading") {
    if (status == "authenticated") {
      router.push("/");
      return;
    } else if (status == "unauthenticated") {
      return children;
    }
  }
}

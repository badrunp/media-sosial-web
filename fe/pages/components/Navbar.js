import { signOut, useSession } from "next-auth/react";
import Link from "next/link";
import { types, useStore } from "../_app";

export default function Navbar() {
  const { data: session } = useSession();
  const { dispatch } = useStore();

  const handleSignout = () => {
    dispatch({ type: types.SET_USER, payload: { user: null } });
    signOut({ redirect: false });
  };
  
  return (
    <div className="flex justify-between items-center bg-gray-700 fixed top-0 left-0 w-full">
      <div className="flex">
        <Link href={"/"}>
          <a className="block text-white py-2 text-sm px-4">Home</a>
        </Link>
        <Link href={"/about"}>
          <a className="block text-white py-2 text-sm px-4">About</a>
        </Link>
        <Link href={"/signin"}>
          <a className="block text-white py-2 text-sm px-4">Signin</a>
        </Link>
      </div>
      <div className="flex">
        {session?.user && (
          <button
            onClick={handleSignout}
            className="block text-white py-2 text-sm px-4"
          >
            Logout
          </button>
        )}
      </div>
    </div>
  );
}

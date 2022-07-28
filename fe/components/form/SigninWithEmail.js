import Image from "next/image";
import React from "react";

function SigninWithEmail() {
  return (
    <div className="mt-8 border-t border-zinc-400/60 pt-8 relative">
      <span className="absolute -top-3 text-sm left-1/2 -translate-x-1/2 bg-white px-4 font-semibold text-zinc-500 tracking-wide truncate">
        or Sign in with Email
      </span>
      <button
        type="button"
        className="w-full space-x-3 px-6 py-3 border border-zinc-400/70 rounded-full font-semibold text-sm sm:text-base flex items-center justify-center"
      >
        <Image src={"/images/google.png"} width="25" height="25" />
        <span>Sign in with Google</span>
      </button>
    </div>
  );
}

export default SigninWithEmail;

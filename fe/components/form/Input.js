import React from "react";

function Input({ type = "text", name, placeholder = "", onChange }) {
  return (
    <input
      type={type}
      name={name}
      id={name}
      className="tracking-wide border border-zinc-400/60 rounded-full w-full focus:outline-none px-6 py-3 mt-1  placeholder:text-zinc-300 text-sm sm:text-base focus:ring-2 focus:border-violet-700 focus:ring-violet-100 transition-colors ease-in-out"
      autoComplete="off"
      placeholder={placeholder}
      onChange={onChange}
    />
  );
}

export default Input;

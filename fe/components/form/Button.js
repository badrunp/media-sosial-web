import React from "react";

function Button({ type = "button", handleClick, children }) {
  return (
    <button
      type={type}
      className="w-full px-6 py-3 bg-violet-600 rounded-full text-white font-semibold text-sm sm:text-base focus:outline-none focus:ring-2 focus:ring-offset-1 focus:ring-violet-600 transition-colors ease-in-out"
      onClick={handleClick && handleClick}
    >
      {children}
    </button>
  );
}

export default Button;

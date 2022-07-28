import React from "react";

function Checkbox({ name }) {
  return (
    <input
      type="checkbox"
      className="rounded-sm border-zinc-400/70 cursor-pointer text-violet-600 focus:ring-1 focus:ring-offset-1 focus:ring-violet-600 block font-semibold text-sm sm:text-base"
      id={name}
    />
  );
}

export default Checkbox;

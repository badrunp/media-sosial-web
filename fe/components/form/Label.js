import React from "react";

function Label({ name, label = "" }) {
  return (
    <label
      htmlFor={name.toLowerCase()}
      className="block font-semibold text-sm sm:text-base"
    >
      {label}
    </label>
  );
}

export default Label;

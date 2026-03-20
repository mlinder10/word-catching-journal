import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

// color is 6-digit hex, ex: #ff0000
export function isColorLight(color: string) {
  const r = parseInt(color.slice(1, 3), 16);
  const g = parseInt(color.slice(3, 5), 16);
  const b = parseInt(color.slice(5, 7), 16);
  return r * 0.299 + g * 0.587 + b * 0.114 > 186;
}

// color is 6-digit hex, ex: #ff0000
// if color is light, shift it slightly darker
// if color is dark, shift it slightly lighter
export function shiftColor(color: string) {
  const r = parseInt(color.slice(1, 3), 16);
  const g = parseInt(color.slice(3, 5), 16);
  const b = parseInt(color.slice(5, 7), 16);
  const factor = isColorLight(color) ? 0.9 : 1.1;
  return `#${Math.trunc(r * factor)
    .toString(16)
    .padStart(2, "0")}${Math.trunc(g * factor)
    .toString(16)
    .padStart(2, "0")}${Math.trunc(b * factor)
    .toString(16)
    .padStart(2, "0")}`;
}

export function isAlpha(str: string) {
  // match a-z, A-Z, and hyphens
  return /^[a-zA-Z-]+$/.test(str);
}

export function capitalize(str: string) {
  if (!str.length) return str;
  if (str.length === 1) return str.toUpperCase();
  return str[0].toUpperCase() + str.slice(1);
}

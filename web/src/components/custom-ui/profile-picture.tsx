import { cn, isColorLight, shiftColor } from "@/lib/utils";
import { useState } from "react";

type ProfilePictureProps = {
  username: string;
  imageUrl: string | null;
  color: string;
  size?: number;
};

export default function ProfilePicture({
  username,
  imageUrl,
  color,
  size = 48,
}: ProfilePictureProps) {
  const [imageError, setImageError] = useState(false);
  const isLight = isColorLight(color);
  const shifted = shiftColor(color);

  if (imageError || !imageUrl) {
    return (
      <div
        className={cn(
          "place-items-center grid border rounded-full",
          isLight ? "text-black" : "text-white",
        )}
        style={{
          backgroundColor: color,
          width: size,
          height: size,
          border: `2px solid ${shifted}`,
          fontSize: size / 2,
        }}
      >
        {(username.at(0) ?? "?").toUpperCase()}
      </div>
    );
  }
  return (
    <img
      src={imageUrl}
      className="rounded-full"
      style={{ width: size, height: size }}
      onError={() => setImageError(true)}
    />
  );
}

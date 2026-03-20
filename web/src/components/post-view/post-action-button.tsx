import { Bookmark, Heart, type LucideIcon } from "lucide-react";
import { Button } from "../ui/button";
import api from "@/api";
import { usePostAction } from "./use-post-action";

type PostActionButtonProps = {
  icon: LucideIcon;
  count: number;
  active: boolean;
  activeColor: string;
  onClick: () => void;
  disabled?: boolean;
};

function PostActionButton({
  icon: Icon,
  count,
  active,
  activeColor,
  onClick,
  disabled,
}: PostActionButtonProps) {
  return (
    <Button variant="ghost" onClick={onClick} disabled={disabled} size="sm">
      <Icon
        fill={active ? "currentColor" : "none"}
        style={{ color: activeColor }}
      />
      <span className="ml-2">{count}</span>
    </Button>
  );
}

type ActionButtonProps = {
  postId: string;
  count: number;
  isMarked: boolean;
};

export function PostLikeButton({ postId, count, isMarked }: ActionButtonProps) {
  const {
    handleToggle: handleLike,
    isMarked: isLiked,
    isPending,
    displayCount,
  } = usePostAction(postId, isMarked, count, (postId, isBookmarked) =>
    api.toggleLikePost(postId, isBookmarked),
  );

  return (
    <PostActionButton
      icon={Heart}
      count={displayCount}
      active={isLiked}
      activeColor={"red"}
      onClick={handleLike}
      disabled={isPending}
    />
  );
}

export function PostBookmarkButton({
  postId,
  count,
  isMarked,
}: ActionButtonProps) {
  const {
    handleToggle: handleBookmark,
    isMarked: isBookmarked,
    isPending,
    displayCount,
  } = usePostAction(postId, isMarked, count, (postId, isBookmarked) =>
    api.toggleBookmarkPost(postId, isBookmarked),
  );

  return (
    <PostActionButton
      icon={Bookmark}
      count={displayCount}
      active={isBookmarked}
      activeColor={"var(--primary)"}
      onClick={handleBookmark}
      disabled={isPending}
    />
  );
}

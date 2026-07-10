import type { Announcement } from "../types"
import { admin, teacher } from "./users"

export const announcements: Announcement[] = [
  {
    id: "ann-1",
    title: "Mid-semester exam schedule released",
    body: "Please review the mid-semester exam schedule on your dashboard. Good luck!",
    authorId: admin.id,
    createdAt: new Date("2024-07-09T10:00:00Z"),
  },
  {
    id: "ann-2",
    title: "New assignment: Algebra Basics Quiz",
    body: "The quiz is now open and due this Friday. Don't forget to review module 1.",
    authorId: teacher.id,
    courseId: "course-1",
    createdAt: new Date("2024-07-08T14:30:00Z"),
  },
  {
    id: "ann-3",
    title: "Parent-teacher meeting next week",
    body: "We invite all parents to join the online parent-teacher meeting on Tuesday.",
    authorId: teacher.id,
    createdAt: new Date("2024-07-07T09:00:00Z"),
  },
  {
    id: "ann-4",
    title: "System maintenance",
    body: "Plato will be under maintenance this Sunday from 02:00 to 04:00 WIB.",
    authorId: admin.id,
    createdAt: new Date("2024-07-06T16:00:00Z"),
  },
]

export function getRecentAnnouncements(limit = 5): Announcement[] {
  return [...announcements]
    .sort((a, b) => b.createdAt.getTime() - a.createdAt.getTime())
    .slice(0, limit)
}

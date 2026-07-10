import type { ScheduleEvent } from "../types"

const today = new Date()
today.setHours(0, 0, 0, 0)

function atTime(hour: number, minute: number, dayOffset = 0): Date {
  const d = new Date(today)
  d.setDate(d.getDate() + dayOffset)
  d.setHours(hour, minute, 0, 0)
  return d
}

export const scheduleEvents: ScheduleEvent[] = [
  {
    id: "sched-1",
    courseId: "course-1",
    moduleId: "mod-math-2",
    title: "Live Class: Linear Equations",
    startAt: atTime(8, 0),
    endAt: atTime(9, 30),
    type: "live-class",
  },
  {
    id: "sched-2",
    courseId: "course-2",
    moduleId: "mod-phys-1",
    title: "Lab: Motion and Forces",
    startAt: atTime(10, 0),
    endAt: atTime(11, 30),
    type: "live-class",
  },
  {
    id: "sched-3",
    courseId: "course-1",
    moduleId: "mod-math-1",
    title: "Quiz: Algebra Basics due",
    startAt: atTime(23, 59),
    endAt: atTime(23, 59),
    type: "assignment-due",
  },
  {
    id: "sched-4",
    courseId: "course-3",
    moduleId: "mod-bi-2",
    title: "Reading: Membaca Naratif",
    startAt: atTime(13, 0),
    endAt: atTime(14, 30),
    type: "live-class",
  },
  {
    id: "sched-5",
    courseId: "course-4",
    moduleId: "mod-eng-1",
    title: "Exercise: Present Tenses due",
    startAt: atTime(1, 0),
    endAt: atTime(1, 0),
    type: "assignment-due",
  },
  {
    id: "sched-6",
    courseId: "course-2",
    moduleId: "mod-phys-2",
    title: "Mid-semester: Physics",
    startAt: atTime(8, 0, 3),
    endAt: atTime(10, 0, 3),
    type: "exam",
  },
]

export function getScheduleForCourseIds(courseIds: string[]): ScheduleEvent[] {
  return scheduleEvents
    .filter((e) => courseIds.includes(e.courseId))
    .sort((a, b) => a.startAt.getTime() - b.startAt.getTime())
}

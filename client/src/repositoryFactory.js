import ColorRepository from './colorRepository'
import TimeRepository from './timeRepository'

const repositories = {
  colors: ColorRepository,
  time: TimeRepository
}

export const RepositoryFactory = {
  get: name => repositories[name]
}

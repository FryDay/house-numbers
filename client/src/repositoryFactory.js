import ColorRepository from './colorRepository'
import TimeRepository from './timeRepository'
import ForceRepository from './forceRepository'

const repositories = {
  color: ColorRepository,
  time: TimeRepository,
  force: ForceRepository
}

export const RepositoryFactory = {
  get: name => repositories[name]
}

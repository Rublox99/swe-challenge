import type { Hit, ZincResponse } from '@/interfaces/Zinc'
import axios from 'axios'

class ZincService {
    private static baseUrl = "http://localhost:8080/emails"

    static async GetFilteredEmails(from: number, size: number, startDate: string, endDate: string, text: string): Promise<ZincResponse> {
        try {
            const response = await axios.get<ZincResponse>(`${this.baseUrl}/filters`, {
                params: {
                    from, size,
                    start_date: startDate,
                    end_date: endDate,
                    text
                }
            })

            return response.data
        } catch (error) {
            throw new Error(`Failed to fetch the filtered emails: ${error}`)
        }
    }

    static async GetAllEmails(from: number, size: number): Promise<ZincResponse> {
        try {
            const response = await axios.get<ZincResponse>(`${this.baseUrl}/all`, {
                params: {
                    from, size
                }
            })

            return response.data
        } catch (error) {
            throw new Error(`Failed to fetch all the emails: ${error}`)
        }
    }

    static async GetEmailById(id: string): Promise<Hit> {
        try {
            const response = await axios.get<Hit>(`${this.baseUrl}/index`, {
                params: {
                    id
                }
            })

            return response.data
        } catch (error) {
            throw new Error(`Failed to fetch the email by ID: ${error}`)
        }
    }
}

export {
    ZincService
}